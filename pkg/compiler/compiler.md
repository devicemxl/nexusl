## Fase 3: Compilación a Go como Lenguaje Intermedio

Objetivo: Transformar el AST de NexusL en código Go, que luego puede ser compilado a binarios nativos.

### pkg/compiler/compiler.go (MODIFICADO/EXPANDIDO)

El rol del compilador sería tomar un ast.Program (el AST completo de NexusL) y generar código Go.
Mapeo de AST a Go:
- FlatTripletaStatement de tipo cli print "hello"; se traduciría a fmt.Println("hello") en Go.
- FlatTripletaStatement de tipo assert Juan tiene-edad 30; se traduciría a una llamada a la API de CozoDB: db.AddTripleta("Juan", "tiene-edad", "30").
- Las S-expressions que definen reglas se traducirían a sentencias db.DefineRule() o se representarían de otra forma que CozoDB pueda consumir.
- Las consultas (con variables) serían las más interesantes: Se traducirían a llamadas a db.Query() y luego el código Go iteraría sobre los resultados.

Proceso de Compilación:

El cmd/nexusl/main.go ya no solo evaluaría, sino que tendría un modo de "compilación".
nexusl compile <input_file.nli> -o <output_executable>

Pasos:
- Leer input_file.nli.
- Lexer -> Parser -> AST.
- compiler.Compile(AST) -> Generar un archivo .go temporal (o varios).
- Ejecutar el compilador de Go: go build -o <output_executable> temp_nexusl_code.go.
- Opcional: Limpiar archivos temporales.