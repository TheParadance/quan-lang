import { useEffect, useState } from 'react'
import './App.css'
import qlangExecute, { loadWasm } from './lib/qlang'
import Editor from 'react-simple-code-editor';
import Prism from 'prismjs/components/prism-core';
import 'prismjs/components/prism-clike';
import 'prismjs/components/prism-javascript';
// import 'prismjs/themes/prism.css'; //Example style, you can use another
import 'prismjs/themes/prism-okaidia.css';
import 'prismjs/components/prism-javascript';

import ReactJson from 'react-json-view';




const sampleCode = `println("Hello world")\nx = 1\ny = 2\nz = x + y\nprintln(z)`

function App() {
  const [code, setCode] = useState(sampleCode)
  const [result, setResult] = useState({
    outputs: {},
    console: '',
    ast: [],
    tokens: []
  })


  useEffect(() => {
    loadWasm('/qlang.wasm')
  }, [])

  const exe = async () => {
    const r = await qlangExecute(code, {})
    console.log("r", r)
    setResult((r as any).payload);
  }


  return (
    <div className='flex flex-col w-full h-full md:overflow-hidden overflow-y-auto'>
      <div className='py-3 px-3 flex justify-between items-center shadow'>
        <div className='font-bold text-[1.1rem]'>QuanLang <span className='text-gray-800 font-light text-[0.8rem]'>{`v1.4`}</span></div>
        <button onClick={exe} className='px-3 py-1 rounded-md cursor-pointer bg-black text-white active:scale-[0.95] transition-all duration-150'>Execute</button>
      </div>
      <div className='w-full md:h-full h-max flex flex-col md:flex-row md:overflow-hidden overflow-y-auto'>
        <Editor
          value={code}
          onValueChange={code => setCode(code)}
          highlight={code => Prism.highlight(code, Prism.languages.js)}
          padding={10}
          className='md:w-[35%] w-full md:h-full min-h-[50vh] outline-none'
          style={{
            backgroundColor: '#1e1e1e',
            color: '#d4d4d4',
            fontFamily: '"Fira code", "Fira Mono", monospace',
            fontSize: 16,
            outline: 'none'
          }}
        />
        <div className='w-full md:w-[65%] md:h-full h-max md:overflow-hidden flex flex-col md:border-l-[10px] md:border-l-gray-600 md:border-t-0 order-t-[10px] border-t-gray-600'>
          <div className='w-full md:min-h-[50%] min-h-[100vh] overflow-y-auto bg-gray-950 whitespace-pre-wrap text-white p-5' style={{ fontFamily: '"Fira code", "Fira Mono", monospace', }}>
            {`${result.console || ''}`}
          </div>
          <div className='w-full md:h-[50%] h-max overflow-hidden flex md:flex-row flex-col bg-gray-800 whitespace-pre-wrap text-white gap-2'>
            <div className='flex flex-col md:w-[50%] w-full overflow-hidden'>
              <div className='bg-white text-black py-2 px-1 font-bold'>Tokens</div>
              <ReactJson style={{ overflowY: 'auto' }} src={result.tokens} theme="monokai" />
            </div>
            <div className='flex flex-col md:w-[50%] w-full overflow-hidden'>
              <div className='bg-white text-black py-2 px-1 font-bold'>Abstract Syntax Tree</div>
              <ReactJson style={{ overflowY: 'auto' }} src={result.ast} theme="monokai" />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App





