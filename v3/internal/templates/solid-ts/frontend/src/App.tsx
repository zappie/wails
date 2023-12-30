import { createSignal, onMount } from 'solid-js'
import {Greet} from "../bindings/main/GreetService.js";
import {Events} from "@wailsio/runtime";

function App() {
  const [name, setName] = createSignal('');
  const [result, setResult] = createSignal('Please enter your name below 👇');
  const [time, setTime] = createSignal('Listening for Time event...');

  const doGreet = () => {
    if (!name()) {
      setName('from Go');
    }
    Greet(name()).then((resultValue: string) => {
      setResult(resultValue);
    }).catch((err: any) => {
      console.log(err);
    });
  }

  onMount(() => {
    Events.On('time', (timeValue: any) => {
      setTime(timeValue.data);
    });
  });

  return (
    <div class="container">
      <div>
        <a wml-openURL="https://wails.io">
          <img src="/wails.png" class="logo" alt="Wails logo"/>
        </a>
        <a wml-openURL="https://solidjs.com">
          <img src="/solid.svg" class="logo solid" alt="Solid logo"/>
        </a>
      </div>
      <h1>Wails + Solid</h1>
      <div class="card">
        <div class="result">{result()}</div>
        <div class="input-box">
          <input class="input" value={name()} onInput={(e) => setName(e.currentTarget.value)} type="text" autocomplete="off"/>
          <button class="btn" onClick={doGreet}>Greet</button>
        </div>
      </div>
      <div class="footer">
        <div><p>Click on the Wails logo to learn more</p></div>
        <div><p>{time()}</p></div>
      </div>
    </div>
  )
}

export default App