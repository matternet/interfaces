// restart-monitor.js
const {spawn} = require('child_process')
const nodeexe = 'node'
const timer = ms => new Promise( res => setTimeout(res, ms))
timer(5000).then(_ => {
    const sp = spawn(nodeexe, ['monitor.js'], {
        detached: true,
        stdio: 'ignore'
    })
    sp.unref()
    console.log('spawn done')
})
