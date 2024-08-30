package benchmarkgojson

import (
	"bytes"
	"encoding/json"
)

// https://github.com/npm/registry/blob/main/docs/responses/package-metadata.md#abbreviated-metadata-format

type AbbreviatedMetadata struct {
	Name     string                    `json:"name"`
	Modified string                    `json:"modified,omitempty"`
	Versions map[string]*VersionObject `json:"versions"`
	DistTags map[string]string         `json:"dist-tags"`
}

func (o *AbbreviatedMetadata) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer(nil)

	// IMPLEMENT ME

	return b.Bytes(), nil
}

type VersionObject struct {
	Name    string      `json:"name"`
	Version string      `json:"version"`
	Dist    *DistObject `json:"dist"`

	Deprecated       string            `json:"deprecated,omitempty"`
	Dependencies     map[string]string `json:"dependencies,omitempty"`
	DevDependencies  map[string]string `json:"devDependencies,omitempty"`
	PeerDependencies map[string]string `json:"peerDependencies,omitempty"`
}

func (o *VersionObject) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer(nil)
	b.WriteString(`{"name":"`)
	b.WriteString(o.Name)
	b.WriteString(`","version":"`)
	b.WriteString(o.Version)
	b.WriteString(`","dist":`)
	by, err := o.Dist.MarshalJSON()
	if err != nil {
		return nil, err
	}
	b.Write(by)
	if o.Deprecated != "" {
		b.WriteString(`,"deprecated":"`)
		b.WriteString(o.Deprecated)
		b.WriteString(`"`)
	}
	if len(o.Dependencies) > 0 {
		b.WriteString(`,"dependencies":`)
		by, err := json.Marshal(o.Dependencies)
		if err != nil {
			return nil, err
		}
		b.Write(by)
	}
	if len(o.DevDependencies) > 0 {
		b.WriteString(`,"devDependencies":`)
		by, err := json.Marshal(o.DevDependencies)
		if err != nil {
			return nil, err
		}
		b.Write(by)
	}
	if len(o.PeerDependencies) > 0 {
		b.WriteString(`,"peerDependencies":`)
		by, err := json.Marshal(o.PeerDependencies)
		if err != nil {
			return nil, err
		}
		b.Write(by)
	}
	b.WriteString("}")
	return b.Bytes(), nil
}

type DistObject struct {
	Tarball string `json:"tarball,omitempty"`
	ShaSum  string `json:"shasum,omitempty"`
}

func (o *DistObject) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer(nil)
	b.Grow(len(o.Tarball) + len(o.ShaSum) + 26)
	b.WriteString(`{"tarball":"`)
	b.WriteString(o.Tarball)
	b.WriteString(`","shasum":"`)
	b.WriteString(o.ShaSum)
	b.WriteString(`"}`)
	return b.Bytes(), nil
}
