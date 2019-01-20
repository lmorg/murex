// +build windows

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_windows.mx using docgen.

   Please do not manually edit this file because it will be overwritten.
*/

func init() {
	murexProfile = append(murexProfile, "autocomplete set go { [{\n    \"Flags\": [ \"build\", \"clean\", \"doc\", \"env\", \"bug\", \"fix\", \"fmt\", \"generate\", \"get\", \"install\", \"list\", \"run\", \"test\", \"tool\", \"version\", \"vet\", \"help\" ]\n}] }\n\n\nautocomplete set cast { [{\n    \"Dynamic\": ({ runtime: --unmarshallers })\n}] }\n\nautocomplete set tout { [{\n    \"Dynamic\": ({ runtime: --marshallers })\n}] }")
}
