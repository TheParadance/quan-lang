
export async function loadWasm(fileUrl: string) {
    const go = new (window as any).Go(); // `Go` is injected by wasm_exec.js

    const result = await WebAssembly.instantiateStreaming(
        fetch(fileUrl),
        go.importObject
    );

    await go.run(result.instance);
}


export default async function qlangExecute(program: string, vars: any) {
    const result = await (window as any).execute({
        mode: "RELEASE",
        program: program,
        vars: vars,
        debugLv: [],
    });
    try {
        result.payload.ast = JSON.parse(result.payload.ast);
        result.payload.tokens = JSON.parse(result.payload.tokens)
    } catch (e) {
        console.error('fail to parse', e)
    } finally {
        console.log("Result from Go:", result);
    }
    return result
}