package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value string
	IsSet bool
}

func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %s, IsSet: %t", c.Key, c.Value, c.IsSet)
}

/*
%v - deaults to the default format for the type
%+v - adds field names when printing structs
%#v - prints the Go syntax representation of the value
%T - prints the type of the value
%% - prints a literal percent sign
%s - formats a string
%q - formats a string as a double-quoted string
%t - formats a boolean value
%d - formats an integer in base 10
%b - formats an integer in base 2
%f - formats a floating-point number
*/
func main() {
	appName := "EnvParser"
	version := 1.0
	port := 8080
	isEnabled := true

	status := fmt.Sprintf("Application: %s, Version: %.1f, Port: %d, Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)

	item1 := ConfigItem{Key: "API_KEY", Value: "http://localhost:3000/api", IsSet: true}
	item2 := ConfigItem{Key: "DB_HOST", Value: "5000", IsSet: true}
	item3 := ConfigItem{Key: "DEBUG_MODE", Value: "false", IsSet: false}

	fmt.Printf("Config Item 1: %v\n", item1)
	fmt.Printf("Config Item 2: %+v\n", item2)
	fmt.Printf("Config Item 3: %#v\n", item3)

}
