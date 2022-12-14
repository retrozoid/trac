package trac

// Represents a single name-value relation.
type Entry struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Represents a list of Entry entities.
type Entries struct {
	Count int32   `json:"count,omitempty"`
	Entry []Entry `json:"entry,omitempty"`
}

// Represents a href link.
type Href struct {
	Href string `json:"href,omitempty"`
}

// Represents a list of MetaData entities.
type Datas struct {
	Count int32      `json:"count,omitempty"`
	Data  []MetaData `json:"data,omitempty"`
}

// Represents a list of items (strings).
type Items struct {
	Item []string `json:"item,omitempty"`
}

// Represents a list of URLs.
type Link struct {
	Type        string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
	RelativeUrl string `json:"relativeUrl,omitempty"`
}

// Represents a list of Link entities.
type Links struct {
	Count int32  `json:"count,omitempty"`
	Link  []Link `json:"link,omitempty"`
}

// Represents a named Entries entity.
type MetaData struct {
	Id      string   `json:"id,omitempty"`
	Entries *Entries `json:"entries,omitempty"`
}

// Represents a list of OperationResult entities.
type MultipleOperationResult struct {
	Count           int32             `json:"count,omitempty"`
	ErrorCount      int32             `json:"errorCount,omitempty"`
	OperationResult []OperationResult `json:"operationResult,omitempty"`
}

// Represents a relation between a message and a related entity.
type OperationResult struct {
	Message string         `json:"message,omitempty"`
	Related *RelatedEntity `json:"related,omitempty"`
}

func NewProperty(name, value string) Property {
	return Property{Name: name, Value: value}
}

// Represents a name-value-type relation.
type Property struct {
	Name      string     `json:"name,omitempty"`
	Value     string     `json:"value"`
	Inherited bool       `json:"inherited,omitempty"`
	Type      *ModelType `json:"type,omitempty"`
}

// Represents a list of RelatedEntity entities.
type RelatedEntities struct {
	Count  int32           `json:"count,omitempty"`
	Entity []RelatedEntity `json:"entity,omitempty"`
}

// Represents a related entity.
type RelatedEntity struct {
	Type       string     `json:"type,omitempty"`
	Unknown    bool       `json:"unknown,omitempty"`
	InternalId string     `json:"internalId,omitempty"`
	Text       string     `json:"text,omitempty"`
	Build      *Build     `json:"build,omitempty"`
	BuildType  *BuildType `json:"buildType,omitempty"`
	Project    *Project   `json:"project,omitempty"`
	User       *User      `json:"user,omitempty"`
	Group      *Group     `json:"group,omitempty"`
	Test       *Test      `json:"test,omitempty"`
	Problem    *Problem   `json:"problem,omitempty"`
	Agent      *Agent     `json:"agent,omitempty"`
	VcsRoot    *VcsRoot   `json:"vcsRoot,omitempty"`
	Change     *Change    `json:"change,omitempty"`
	AgentPool  *AgentPool `json:"agentPool,omitempty"`
}

// Represents a list of Property entities.
type Properties struct {
	Count    int32      `json:"count,omitempty"`
	Href     string     `json:"href,omitempty"`
	Property []Property `json:"property,omitempty"`
}

func (p *Properties) Add(v ...Property) {
	p.Property = append(p.Property, v...)
	p.Count += int32(len(v))
}

func (p *Properties) Get(name string) *Property {
	for n := range p.Property {
		if p.Property[n].Name == name {
			return &p.Property[n]
		}
	}
	return nil
}

// Represents a link to the Builds entity.
type Related struct {
	Builds *Builds `json:"builds,omitempty"`
}

// Represents a build parameter type string.
type ModelType struct {
	RawValue string `json:"rawValue,omitempty"`
}

// Represents a dated comment of the specific user.
type Comment struct {
	Timestamp string `json:"timestamp,omitempty"`
	Text      string `json:"text,omitempty"`
	User      *User  `json:"user,omitempty"`
}
