import { InputHTMLAttributes, useEffect, useState, } from 'react'

import {
  setListener,
  useButton,
  useNumber,
  useObject,
  useString
} from "@pathserve/messenger";
function App() {


  // input 1
  const [valueInput1, setInput1] = useState("input 1")
  const [sendInput1] = useString("input 1",  valueInput1, (newValue) => {
    setInput1(newValue)
  })

  // input 2
  const [valueInput2, setInput2] = useState(0)
  const [sendInput2] = useNumber("input 2",valueInput2 , (newValue)=>{
    setInput2(newValue)
  })
  // input 3
  const [valueInput3, setInput3] = useState({id: 1, text: "txt"})
  const [sendInput3] = useObject("input 3",valueInput3 , (newValue)=>{
    // @ts-expect-error: ""
    setInput3(newValue)
  })
  // input 4
  const [ count, setCount] = useState(0)
  useButton("button", () => {
    setCount(count +1)
  })

  useEffect(() => {
    setListener()
  }, [])
  return (
    <>
    <h1>React</h1>
    <div>
        <p>input 1</p>
        <input type="string" onChange={(event: React.FormEvent<HTMLInputElement>) => {sendInput1(event.currentTarget.value)}} value={valueInput1}/>
      </div>
      <div>
        <p>input 2</p>
        <input type="number" onChange={(event: React.FormEvent<HTMLInputElement>) => {sendInput2(parseInt(event.currentTarget.value))}} value={valueInput2}/>
      </div>
      <div>
        <p>input 3</p>
        <pre>
          {JSON.stringify(valueInput3)}
        </pre>
        <button onClick={() => {sendInput3({...valueInput3, id: valueInput3.id + 1})}}>change object</button>
      </div>
      <div>
        <p>button</p>
        <p>{count}</p>
      </div>
    </>
  )
}

export default App
