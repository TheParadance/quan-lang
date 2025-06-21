async function main() {
    const res = await fetch('http://localhost:3000/execute', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            program: `
                fn isGraterthen(x, y){
                   if (x > y){
                       return true;
                   } else {
                       return false;
                   }
                }

                fn main(x, y){
                    if(isGraterthen(x, y)){
                        return 1;
                    } else {
                        return 2;
                    }
                }
                flag = main(x, y)
            `,
            vars: {
                x: 5,
                y: 9
            },
        })
    })
    const data = await res.json()
    console.log(data)
}
main()