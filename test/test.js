async function main() {
    const res = await fetch('http://localhost:3000/execute', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            program: `
                print('''hello world \${x}''')
                print('''hello world \${x}''')
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