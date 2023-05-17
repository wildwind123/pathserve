import {useString, setListener, useButton} from "@pathserve/messenger"


export const html = `
<div>
    <h1>Vanilla</h1>
    <p>input 1</p>
    <input id="input1" />
    <p>button event</p>
    <div id="button-count">0</div>
</div>
`



export const script = () => {
    document.addEventListener("DOMContentLoaded", function(){
        setListener()
    });
    // input 1
    const input1El = document.getElementById("input1") as HTMLInputElement

    const [setInput1] = useString("input 1", input1El.value, (newValue) => {
        input1El.value = newValue
    } )
    input1El.addEventListener('input', (e: Event) => {
        setInput1((e.target as HTMLInputElement).value)
    })
    // button
    const buttonCountEl = document.getElementById("button-count") as HTMLDivElement
    useButton('button', () => {
        buttonCountEl.textContent = (parseInt(buttonCountEl.textContent ?? '0') +1).toString()
    })
}