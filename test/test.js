async function main() {
    const res = await fetch('http://localhost:3000/execute', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            program: `
                a = {
                    x: "hello",
                    y: 20,
                    z: {
                        w: 40
                    }
                }
                print('''hello world \${a.x}''')
                a.z.w = 100
                print('''hello world \${a.z.w}''')
            `,
            vars: {
                x: 10
            },
        })
    })
    const data = await res.json()
    console.log(data)
}
main()