
from hashlib import sha256 as sha
from json import dumps
#
def generate_chunk_id(content: str) -> str:
    """Genera un ID único para el chunk usando un hash SHA256."""
    return sha((content).encode('utf-8')).hexdigest()
#
def parse_chunks(input_filepath: str, output_filepath: str):
    allChunks       = {}
    chunk_counter   = 0
    processed_chunk = None
    content_data    = None
    #
    with open(input_filepath, 'r', encoding='utf-8') as f:
        lines = f.readlines() # Leer todas las líneas de una vez para un mejor manejo de bloques
    for line in lines:
        #
        # Procesamiento de campos
        #
        if "ID_AUTO" in line:
            continue
        elif "---CHUNK_START---" in line:
            processed_chunk = dict()
        elif (processed_chunk != None) and "source_file:" in line:
            thisData = "".join(line.split(":")[1]).replace('\n','').replace('\t','').lstrip()
            processed_chunk["source_file"] = thisData
        elif (processed_chunk != None) and "section_heading:" in line:
            thisData = "".join(line.split(":")[1]).replace('\n','').replace('\t','').lstrip()
            processed_chunk["section_heading"] = thisData
        elif (processed_chunk != None) and "page_number:" in line:
            thisData = "".join(line.split(":")[1]).replace('\n','').replace('\t','').lstrip()
            processed_chunk["page_number"] = thisData
        elif (processed_chunk != None) and "language:" in line:
            thisData = "".join(line.split(":")[1]).replace('\n','').replace('\t','').lstrip()
            processed_chunk["language"] = thisData
        elif (processed_chunk != None) and "group:" in line:
            thisData = "".join(line.split(":")[1]).replace('\n','').replace('\t','').lstrip()
            processed_chunk["group"] = thisData
        elif (processed_chunk != None) and "tags:" in line:
            thisData = [x.lstrip() for x in "".join(line.split(":")[1]).replace('\n','').replace('\t','').lstrip().split(",")]
            processed_chunk["tags"] = thisData
        #
        # Se inicia procesamiento de contenido multiliean?
        #
        elif "content_start" in line:
            content_data = list()
        #
        elif content_data != None and "content_end" not in line:
            # Aquí, solo añade la línea, sin hacer escapes manuales
            content_data.append(line) # No .replace('"', "&quot; ").replace("'", "&#39; ")
            # content_data.append('\n') # Esto tampoco es necesario si la línea ya viene con un salto de línea y lo quieres conservar
        elif "content_end" in line:
            # Reemplaza 'content_data = ",".join([x.lstrip() for x in content_data])'
            # Simplemente une las líneas sin comas adicionales, y elimina los saltos de línea extra si los hay al final.
            processed_chunk["content"] = "".join(content_data).strip() 
            # Asegúrate de que no haya comas extra o saltos de línea dobles al unir.
            # Puedes probar: processed_chunk["content"] = "\n".join([l.rstrip('\n') for l in content_data]).strip()
            # Esto eliminará saltos de línea al final de cada línea antes de unirlas con '\n'
            
            content_data = None
        #
        # Si se detecta finalizacion de chunk
        #
        elif "---CHUNK_END---" in line:
            #allData = pd.concat([allData, pd.DataFrame(processed_chunk)], ignore_index=True)
            idx = generate_chunk_id(str(processed_chunk))
            processed_chunk["id"] = idx
            allChunks[idx] = processed_chunk
            
    #
    # JSON outoput
    #
        # Esto es crucial: usar json.dumps() con el diccionario Python directamente.
    # NO: f.write(str(allChunks).replace("'", '"') + '\n')
    # SÍ:
    with open(output_filepath, 'w', encoding='utf-8') as f:
        # Itera sobre los chunks en el diccionario allChunks
        #f.write(dumps(allChunks, ensure_ascii=False) + '\n') 
        for chunkNum in allChunks.keys():
            # json.dumps() se encarga automáticamente de escapar las comillas dobles internas
            # y de manejar los caracteres Unicode.
            print(allChunks[chunkNum])
            f.write(dumps(allChunks[chunkNum], ensure_ascii=False) + '\n')
    print(f"Chunks procesados y guardados en {output_filepath}")
#
parse_chunks("lab/sourceOfTrue.txt", "lab/processed_chunks_safe.jsonl")
