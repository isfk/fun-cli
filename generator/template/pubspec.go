package template

var PubspecTPL = `// Generated By isfk/flutter-cli
name: {{.AppName}}
description: A new Flutter With Get project.
version: 1.0.0+1

environment:
  sdk: ">=2.19.2 <3.0.0"

dependencies:
  flutter:
    sdk: flutter
  cupertino_icons: ^1.0.2
  get: ^4.6.5
  get_storage: ^2.0.3

dev_dependencies:
  flutter_test:
    sdk: flutter
  flutter_lints: ^2.0.0

flutter:
  uses-material-design: true
`
