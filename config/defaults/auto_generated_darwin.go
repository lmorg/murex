// +build darwin

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_osx.mx using docgen.

   Please do not manually edit this file because it will be automatically
   overwritten by the build pipeline. Instead please edit the aforementioned
   profile_osx.mx file located in the same directory.
*/

func init() {
	murexProfile = append(murexProfile, "autocomplete set brew { [\n\t{\n\t\t\"Flags\": [\n            # Built-in commands\n            \"--cache\", \"--env\", \"--version\", \"cleanup\", \"deps\", \"fetch\", \"home\", \"leaves\", \"log\", \"options\", \"postinstall\", \"search\", \"tap-info\", \"unlink\", \"update-report\", \"upgrade\", \"--caskroom\", \"--prefix\", \"analytics\", \"commands\", \"desc\", \"gist-logs\", \"info\", \"link\", \"migrate\", \"outdated\", \"readall\", \"shellenv\", \"tap\", \"unpin\", \"update-reset\", \"uses\", \"--cellar\", \"--repository\", \"cask\", \"config\", \"doctor\", \"help\", \"install\", \"list\", \"missing\", \"pin\", \"reinstall\", \"switch\", \"uninstall\", \"untap\", \"update\", \"vendor-install\",\n            \n            # Built-in developer commands\n            \"audit\", \"bump-revision\", \"create\", \"extract\", \"linkage\", \"pr-automerge\", \"prof\", \"sh\", \"test\", \"update-license-data\", \"bottle\", \"bump\", \"dispatch-build-bottle\", \"formula\", \"livecheck\", \"pr-publish\", \"pull\", \"sponsors\", \"tests\", \"update-python-resources\", \"bump-cask-pr\", \"cat\", \"diy\", \"install-bundler-gems\", \"man\", \"pr-pull\", \"release-notes\", \"style\", \"typecheck\", \"update-test\", \"bump-formula-pr\", \"command\", \"edit\", \"irb\", \"mirror\", \"pr-upload\", \"ruby\", \"tap-new\", \"unpack\", \"vendor-gems\",\n            \n            # External commands\n            \"aspell-dictionaries\", \"postgresql-upgrade-database\"\n        ],\n        \"FlagValues\": {\n            \"cask\": [{\n                \"Flags\": [\n                    # Cask commands\n                    \"--cache\", \"_help\", \"_stanza\", \"audit\", \"cat\", \"create\", \"doctor\", \"edit\", \"fetch\", \"help\", \"home\", \"info\", \"install\", \"list\", \"outdated\", \"reinstall\", \"style\", \"uninstall\", \"upgrade\", \"zap\",\n\n                    # External cask commands\n                    \"ci\"\n                ]\n            }]\n        }\n\t},\n    {\n        \"DynamicDesc\": ({ \n            brew help @{ $ARGS->@[1..] } -> tabulate: --map --split-comma --column-wraps\n        }),\n        \"ListView\": true\n    }\n] }")
}
