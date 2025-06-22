const fs = require('fs')

async function main() {
    const res = await fetch('http://localhost:3000/execute', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            program: fs.readFileSync('./sample-program.qlang', 'utf-8'),
            vars: {
                "obj": {
                    person: "John",
                    age: 25
                }
            },
        })
    })
    const data = await res.json()
    console.log(data)
}
main()