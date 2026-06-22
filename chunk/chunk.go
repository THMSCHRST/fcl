package chunk

import (
	"fcl/mesh"
	"fcl/td"
	"fmt"
	"math"
	"sync"
)

type loadedChunk struct {
	X, Z int
	Lod  int
	Mesh *mesh.Mesh
}

type ChunkManager struct {
	GenFunc   func(chunkSize, offsetX, offsetZ, lod int, seed int64) *mesh.Mesh
	ViewDist  int
	DistToLod func(int) int
	ChunkSize int
	Seed      int64

	mu      sync.Mutex
	loaded  map[string]*loadedChunk
	lastPos td.Vec3
}

func NewChunkManager(gen func(chunkSize, offsetX, offsetZ, lod int, seed int64) *mesh.Mesh, viewDist int, distToLod func(int) int, chunkSize int, seed int64) *ChunkManager {
	return &ChunkManager{
		GenFunc:   gen,
		ViewDist:  viewDist,
		DistToLod: distToLod,
		ChunkSize: chunkSize,
		Seed:      seed,
		loaded:    make(map[string]*loadedChunk),
		lastPos:   td.NewVec3(0, 0, 0),
	}
}

func (cm *ChunkManager) Update(pos td.Vec3) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	currentX := int(math.Floor(float64(pos.X / float32(cm.ChunkSize))))
	currentZ := int(math.Floor(float64(pos.Z / float32(cm.ChunkSize))))

	wanted := make(map[string]bool)
	for dx := -cm.ViewDist; dx <= cm.ViewDist; dx++ {
		for dz := -cm.ViewDist; dz <= cm.ViewDist; dz++ {
			if max(abs(dx), abs(dz)) > cm.ViewDist {
				continue
			}
			cx := currentX + dx
			cz := currentZ + dz
			key := key(cx, cz)
			wanted[key] = true

			dist := max(abs(dx), abs(dz))
			lod := cm.DistToLod(dist)
			if lod < 1 {
				lod = 1
			}

			if existing, ok := cm.loaded[key]; ok {
				if existing.Lod == lod {
					continue
				}

				existing.Mesh.Destroy()
				delete(cm.loaded, key)
			}
			offsetX := cx * (cm.ChunkSize - 1)
			offsetZ := cz * (cm.ChunkSize - 1)
			m := cm.GenFunc(cm.ChunkSize, offsetX, offsetZ, lod, cm.Seed)
			cm.loaded[key] = &loadedChunk{
				X:    cx,
				Z:    cz,
				Lod:  lod,
				Mesh: m,
			}
		}
	}

	for key, ch := range cm.loaded {
		if !wanted[key] {
			ch.Mesh.Destroy()
			delete(cm.loaded, key)
		}
	}
}

func (cm *ChunkManager) Draw() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for _, ch := range cm.loaded {
		ch.Mesh.Draw()
	}
}

func (cm *ChunkManager) Destroy() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for _, ch := range cm.loaded {
		ch.Mesh.Destroy()
	}
	cm.loaded = make(map[string]*loadedChunk)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func key(x, z int) string {
	return fmt.Sprintf("%d,%d", x, z)
}