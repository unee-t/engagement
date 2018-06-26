// Aiming for:
// https://e{{STAGE}}.unee-t.com?url=https://dev.case.unee-t.com/case/61914&user=233&id=foobar-123&medium=email
const URL = require('url').URL

const url = new URL('https://e.dev.unee-t.com')

const params = new URLSearchParams({
  url: 'https://dev.case.unee-t.com/case/61914',
  user: 21,
  id: 'foobar-' + Math.floor(Math.random() * 1000) + 1,
  medium: 'email'
})
params.sort()
url.search = params

console.log('curl', `"${url.toString()}"`)
