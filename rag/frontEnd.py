import streamlit as st
import requests # You'll need this!

# --- Configuration ---
GO_BACKEND_URL = "http://localhost:8080/query" # Make sure this matches your Go API endpoint

# Function to call your Go API
def get_rag_response(user_query):
    try:
        # Send the query to your Go backend
        response = requests.post(GO_BACKEND_URL, json={"query": user_query})
        response.raise_for_status() # Raise an exception for HTTP errors (4xx or 5xx)

        # Assuming your Go API returns JSON like {"answer": "..."}
        llm_answer = response.json().get("answer", "Error: No answer field in Go API response.")
        return llm_answer
    except requests.exceptions.ConnectionError:
        return "Error: Could not connect to the Go backend. Is it running?"
    except requests.exceptions.RequestException as e:
        return f"An error occurred with the Go API request: {e}"


# PAGE STYLING

#im = Image.open("favicon.ico") # Replace with your icon file
#st.set_page_config(page_title="Your App Title", page_icon=im)
st.markdown(
            """
            <style>
            /*
            
            BARRA STRIMLIT
            
            
            */
            .css-14xtw13 {
                display:none;
                }
            /*
            
            
            FLECHA DE "SEND" EN LA BARRA DE MENSAJES
            ES UN SVG
            
            TRATAR DE QUE CAMBIE DE COLOR CUANDO ENVIA - COMO
            SI FUERA UN BOTON
            
            
            
            */
            .css-9ilocf {
    vertical-align: middle;
    overflow: hidden;
    color: inherit;
    fill: currentcolor;
    display: inline-flex;
    -webkit-box-align: center;
    align-items: center;
    font-size: 1.5rem;
    width: 1.5rem;
    height: 1.5rem;
}
            
            
            /*
            
            INPUT DE TEXTO REDONDEADO
            
            
            */
            .st-gy {
    width: 1514px;
}

.st-gx {
    border-bottom-color: rgb(75, 233, 255);
}

.st-gw {
    border-top-color: rgb(75, 233, 255);
}

.st-gv {
    border-right-color: rgb(75, 233, 255);
}

.st-gu {
    border-left-color: rgb(75, 233, 255);
}
/* 


APP BACKGROUND COLOR


*/
            .css-1dqmzvx {
    position: absolute;
    background: rgb(11, 21, 39);
    color: rgb(250, 250, 250);
    inset: 0px;
    overflow: hidden;
}   
/*


BARRA DE COLORES SUPERIOR
tratar de ahcer efecto de movimineto cuando "piense" el LLM con JS


*/
            .css-1dp5vir {
    position: absolute;
    top: 0px;
    right: 0px;
    left: 0px;
    height: 0.125rem;
    background-image: linear-gradient(90deg, rgb(119 75 255), rgb(128 255 204));
    z-index: 999990;
}

                
            </style>
            """,
            unsafe_allow_html=True,
        )
st.title("NexusL RAG Chatbot")

# Initialize chat history
if "messages" not in st.session_state:
    st.session_state.messages = []

# Display chat messages from history on app rerun
for message in st.session_state.messages:
    with st.chat_message(message["role"]):
        st.markdown(message["content"])

# Accept user input
if prompt := st.chat_input("Ask about NexusL..."):
    # Add user message to chat history and display
    st.session_state.messages.append({"role": "user", "content": prompt})
    with st.chat_message("user"):
        st.markdown(prompt)

    # Display assistant response in chat message container
    with st.chat_message("assistant"):
        with st.spinner("NexusL is thinking..."): # Show a spinner while waiting for Go backend
            # Call your Go API
            response_from_go = get_rag_response(prompt)
            st.markdown(response_from_go) # Display the response

    # Add assistant response to chat history
    st.session_state.messages.append({"role": "assistant", "content": response_from_go})