async function main() {
    const res = await fetch('http://localhost:3000/execute', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            program: `
                fn print(x){
                    return "hello " + x + " i love you"
                }
                flag = print(x)
            `,
            vars: {
                x: " this is a test",
            },
        })
    })
    const data = await res.json()
    console.log(data)
}
main()