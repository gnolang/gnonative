import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {CloseBridge, InitBridge} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Start GnoNativeKit on Wails to see result");
    const updateResultText = (result: string) => setResultText(result);

    function closeBridge() {
        CloseBridge().then(updateResultText);
    }

    function initBridge() {
        InitBridge().then(updateResultText);
    }


    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div  className="input-box">
                <button className="btn" onClick={initBridge}>Init Bridge</button>
                <button className="btn" onClick={closeBridge}>Close Bridge</button>
            </div>
            <div id="result" className="result">{resultText}</div>
        </div>
    )
}

export default App
