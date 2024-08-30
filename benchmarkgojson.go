package benchmarkgojson

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// https://github.com/npm/registry/blob/main/docs/responses/package-metadata.md#abbreviated-metadata-format

type AbbreviatedMetadata struct {
	Name     string                    `json:"name"`
	Modified string                    `json:"modified,omitempty"`
	Versions map[string]*VersionObject `json:"versions"`
	DistTags map[string]string         `json:"dist-tags"`
}

type VersionObject struct {
	Name             string            `json:"name"`
	Version          string            `json:"version"`
	Deprecated       string            `json:"deprecated,omitempty"`
	Dist             *DistObject       `json:"dist"`
	Dependencies     map[string]string `json:"dependencies,omitempty"`
	DevDependencies  map[string]string `json:"devDependencies,omitempty"`
	PeerDependencies map[string]string `json:"peerDependencies,omitempty"`
}

type DistObject struct {
	Tarball string `json:"tarball,omitempty"`
	ShaSum  string `json:"shasum,omitempty"`
}

func writeStringMap(b *bytes.Buffer, m map[string]string) {
	first := true
	for k, v := range m {
		if !first {
			b.WriteString(",")
		}
		first = false
		b.WriteString(`"`)
		b.WriteString(k)
		b.WriteString(`":"`)
		b.WriteString(v)
		b.WriteString(`"`)
	}
}

func (a *AbbreviatedMetadata) MarshalJSON() ([]byte, error) {
	b := bufferPool.Get().(*bytes.Buffer)
	b.Reset()
	defer bufferPool.Put(b)

	b.WriteString(`{"name":"`)
	b.WriteString(a.Name)
	b.WriteString(`","modified":"`)
	b.WriteString(a.Modified)
	b.WriteString(`","versions":{`)

	first := true
	for k, v := range a.Versions {
		if !first {
			b.WriteString(",")
		}
		first = false
		b.WriteString(`"`)
		b.WriteString(k)
		b.WriteString(`":`)
		version, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}
		b.Write(version)
	}
	b.WriteString(`}`)

	if len(a.DistTags) > 0 {
		b.WriteString(`,"dist-tags":{`)
		first = true
		for k, v := range a.DistTags {
			if !first {
				b.WriteString(",")
			}
			first = false
			b.WriteString(`"`)
			b.WriteString(k)
			b.WriteString(`":"`)
			b.WriteString(v)
			b.WriteString(`"`)
		}
		b.WriteString(`}`)
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

func (v *VersionObject) MarshalJSON() ([]byte, error) {
	b := bufferPool.Get().(*bytes.Buffer)
	b.Reset()
	defer bufferPool.Put(b)

	b.WriteString(`{"name":"`)
	b.WriteString(v.Name)
	b.WriteString(`","version":"`)
	b.WriteString(v.Version)
	b.WriteString(`","dist":`)
	dist, err := v.Dist.MarshalJSON()
	if err != nil {
		return nil, err
	}
	b.Write(dist)

	if v.Deprecated != "" {
		b.WriteString(`,"deprecated":"`)
		b.WriteString(v.Deprecated)
		b.WriteString(`"`)
	}
	if len(v.Dependencies) > 0 {
		b.WriteString(`,"dependencies":{`)
		first := true
		for k, v := range v.Dependencies {
			if !first {
				b.WriteString(",")
			}
			first = false
			b.WriteString(`"`)
			b.WriteString(k)
			b.WriteString(`":"`)
			b.WriteString(v)
			b.WriteString(`"`)
		}
		b.WriteString(`}`)
	}
	if len(v.DevDependencies) > 0 {
		b.WriteString(`,"devDependencies":{`)
		first := true
		for k, v := range v.DevDependencies {
			if !first {
				b.WriteString(",")
			}
			first = false
			b.WriteString(`"`)
			b.WriteString(k)
			b.WriteString(`":"`)
			b.WriteString(v)
			b.WriteString(`"`)
		}
		b.WriteString(`}`)
	}
	if len(v.PeerDependencies) > 0 {
		b.WriteString(`,"peerDependencies":{`)
		first := true
		for k, v := range v.PeerDependencies {
			if !first {
				b.WriteString(",")
			}
			first = false
			b.WriteString(`"`)
			b.WriteString(k)
			b.WriteString(`":"`)
			b.WriteString(v)
			b.WriteString(`"`)
		}
		b.WriteString(`}`)
	}
	b.WriteString("}")
	return b.Bytes(), nil
}

func (o *DistObject) MarshalJSON() ([]byte, error) {
	b := bufferPool.Get().(*bytes.Buffer)
	b.Reset()
	defer bufferPool.Put(b)

	b.WriteString(`{"tarball":"`)
	b.WriteString(o.Tarball)
	b.WriteString(`","shasum":"`)
	b.WriteString(o.ShaSum)
	b.WriteString(`"}`)
	return b.Bytes(), nil
}
