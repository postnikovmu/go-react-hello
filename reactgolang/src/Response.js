import {useState, useEffect} from "react";

function Response() {

    const [hello, setHello] = useState([])

    const HelloFromGo = async () => {
        const response = await fetch("http://localhost:8080/hello")
        const hello = await response.json()
        console.log(hello)
        setHello(hello)
      }
      
    useEffect(()=>{HelloFromGo()}, [])
    
    return(
        <div>
            <p>Hello from Response</p>
            <p>Data: {JSON.stringify(hello)}</p>
        </div>
    );
}

export default Response