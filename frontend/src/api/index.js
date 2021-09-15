export const URL = location.protocol + "//" + location.host + "/api"
// export const URL = location.protocol + "//localhost:9000/api"

function MY_FETCH(URL, REQ) {
  return fetch(URL, REQ)
    .then(response => {
      if(!response.ok) {
        throw("HTTP: " + response.status)
      }

      return response.json()
    })
}

export function POST(URL, BODY) {
  const REQ = {
    method: "POST",
    body: JSON.stringify(BODY)
  }
  
  return MY_FETCH(URL, REQ)
}

export function GET(URL, BODY) {
  const REQ = {
    method: "GET",
    body: JSON.stringify(BODY)
  }

  return MY_FETCH(URL, REQ)
}

export function PUT(URL, BODY) {
  const REQ = {
    method: "PUT",
    body: JSON.stringify(BODY)
  }

  return MY_FETCH(URL, REQ)
}

export function PATCH(URL, BODY) {
  const REQ = {
    method: "PATCH",
    body: JSON.stringify(BODY)
  }

  return MY_FETCH(URL, REQ)
}

export function DELETE(URL, BODY) {
  const REQ = {
    method: "DELETE",
    body: JSON.stringify(BODY)
  }

  return MY_FETCH(URL, REQ)
}
