#!/bin/bash

# Define la carpeta raíz
root_dir="."

# Define las carpetas a crear
declare -A directories=(
  ["cmd/interpreter"]="."
  ["cmd/compiler"]="."
  ["internal/core/ast"]="."
  ["internal/core/parser"]="."
  ["internal/core/lexer"]="."
  ["internal/core/evaluator"]="."
  ["internal/core/environment"]="."
  ["internal/core/builtins"]="."
  ["internal/core/bytecode"]="."
  ["internal/commands"]="."
  ["internal/compiler/semantic_analyzer"]="."
  ["internal/compiler/code_generator"]="."
  ["internal/compiler/optimizer"]="."
  ["internal/runtime/bytecode_vm"]="."
  ["internal/runtime/memory"]="."
  ["internal/runtime/native_bindings"]="."
  ["internal/queries"]="."
  ["internal/repositories"]="."
  ["internal/models"]="."
  ["internal/services"]="."
  ["internal/utils"]="."
  ["configs"]="."
  ["examples"]="."
)

# Crea las carpetas
echo "Creando la estructura de directorios..."
for dir in "${!directories[@]}"; do
  echo "Creando directorio: $root_dir/$dir"
  mkdir -p "$root_dir/$dir"
done

# Crea archivos vacíos (opcional, solo para tener los archivos main.go y session_repo.go)
echo "Creando archivos vacíos..."
touch "$root_dir/cmd/interpreter/main.go"
touch "$root_dir/cmd/compiler/main.go"
touch "$root_dir/internal/repositories/session_repo.go"
touch "$root_dir/internal/models/ast.go"
touch "$root_dir/internal/models/object.go"
touch "$root_dir/internal/models/environment.go"
touch "$root_dir/internal/models/bytecode.go"
touch "$root_dir/internal/commands/execute_code.go"
touch "$root_dir/internal/commands/set_variable.go"
touch "$root_dir/internal/commands/load_module.go"
touch "$root_dir/internal/queries/get_variable.go"
touch "$root_dir/internal/queries/get_stack_trace.go"
touch "$root_dir/internal/queries/get_loaded_modules.go"
touch "$root_dir/internal/services/execution_service.go"

echo "Estructura de directorios creada exitosamente."