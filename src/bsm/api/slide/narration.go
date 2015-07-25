package slide

import (
    "fmt"
    "encoding/json"
    "bsm/api/character"
)

// create a new narration slide with an id of 2
func NewNarrationSlide(name string, narrator character.Character, text string) NarrationSlide {
    ns := NarrationSlide {
        id: 2,
        name: name,
        myType: fmt.Sprintf("%T", NarrationSlide{}),
        data: make(map[string]interface{}),
    }
    ns.data["narrator"] = narrator
    ns.data["text"] = text
    return ns
}

type NarrationSlide struct {
    id          int64                       `json:"id"`
    name        string                      `json:"name"`
    myType      string                      `json:"myType"`
    data        map[string]interface{}      `json:"data"`
}

// get the Id
func (ns NarrationSlide) Id() int64 {
    return ns.id
}
 
// get the Name
func (ns NarrationSlide) Name() string {
    return ns.name
}

// get the Type
func (ns NarrationSlide) Type() string {
    return ns.myType
}

// get the Narrator
func (ns NarrationSlide) Narrarator() (character.Character, bool) {
    return getMappedCharacter(ns.data, "narrator")
}

// get the Text
func (ns NarrationSlide) Text() (string, bool) {
    return getMappedString(ns.data, "text")
}

// json.Marshal to {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{} {narrator:bsm.api.character.Character, 
// text:string}} 
func (ns NarrationSlide) MarshalJSON() ([]byte, error) {
    return marshalSlideJSON(ns.id, ns.name, ns.myType, ns.data)
}

// json.Unmarshal from {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func (ns *NarrationSlide) UnmarshalJSON(data []byte) error {
    um := struct{
        Id          int64                       `json:"id"`
        Name        string                      `json:"name"`
        MyType      string                      `json:"myType"`
        Data        map[string]interface{}      `json:"data"`
    }{}
    err := json.Unmarshal(data, &um)
    if err != nil {
        return fmt.Errorf("Unable to unmarshal data: %s", err)
    }
    
    ns.id = um.Id
    ns.name = um.Name
    ns.myType = um.MyType
    ns.data = um.Data
    
    return nil
}
