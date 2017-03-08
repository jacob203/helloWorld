# only public fields can be marsh
```
type testStrcut struct{
    url string `json:"url"`
}

test := testStruct{
    url:"helloWorld",
}

fmt.Printf("%+v\n", test)---1
testBytes, _ := json.Marshal(test)
fmt.Println(string(testBytes))----2

{longUrl:helloWorld}//1
{}//2
```
both fmt and json use reflect to print field values, but json only print the exported fields

# customize Marshal and UnMarshal
Marshal is to call the interface MarshJson, so implementing the function can do a customize Marsh.
```
type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}
func main() {
	_ = json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "Ken", time.Now()},
	)
}

result:
{"id":1,"name":"Ken","lastSeen":"2009-11-10T23:00:00Z"}

func (u *MyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		LastSeen int64  `json:"lastSeen"`
	}{
		ID:       u.ID,
		Name:     u.Name,
		LastSeen: u.LastSeen.Unix(),
	})
}
```
the upper method is to change the format of LastSeen,
as the original struct's lastSeen is time while the new struct's lastSeen is int64.
it works when there are only little fields which need to be changed, but if there are a lot,
it can be cumbersome.

so we create a new struct to inherit the original, and add the changed fields in the new struct, then they will override the old one
```
func (u *MyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*MyUser
	}{
		LastSeen: u.LastSeen.Unix(),
		MyUser:   u,
	})
}
```
it will cause an infinite loop, as Marshing MyUser calls itself. so we use a type define
```
func (u *MyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*MyUser
	}{
		LastSeen: u.LastSeen.Unix(),
		MyUser:   u,
	})
}
```
it shows that type means another struct which has the same as the original.

The same technique can be used for implementing an UnmarshalJSON method.
```
func (u *MyUser) UnmarshalJSON(data []byte) error {
	type Alias MyUser
	aux := &struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}
```
##### Json Marshal
* []byte: marshalled by bytes
* string: used directly
* jsonrawmessage must be a json Object
in json, it can be Json int, Json String, Json Struct, Json float, Json Array, Json Map, Json bool, Json slice
Json int and Json float is a number: it is encoded like {"number":10}
Json string is a string: it is encoded like {"name":"hello"}
Json Array and Json slice starts [], but Json []byte is encoded into base64
Json struct and Json map it starts like {}
so json raw message value must contain a json value, it means its format must like one of json object.
actually all customized object must be one of json objects

