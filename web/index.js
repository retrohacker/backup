const choo = require('choo')
const html = require('choo/html')
const app = choo()

function setupCard(content) {
  return html`
    <body>
      <div class="absolute w-100 h-100 setup-splash-bg"></div>
      <div class="absolute w-75 h-75 br4 bg-white o-80 center"></div>
      <div class="absolute maxw-75 maxh-75 scroll center tc">
        ${content}
      </div>
    </body>
  `
}

app.route('/setup', function () {
  return setupCard(html`
      <div>
        <h1>Backup</h1>
        <h1 class="f1"><i class="fa fa-lock"></i></h1>
        <h3>Ensure your files are safe</h3>
        <div class="f4 link dim br3 ph3 pv3 mb2 dib white bg-yellow">LET'S GET STARTED</div>
      </div>
  `)
})

app.route('*', function (state, emit) {
  emit('pushState', '/setup')
})

app.mount('body')
