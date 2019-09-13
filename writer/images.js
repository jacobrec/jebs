function loadImages() {
  let ls = JSON.parse(files).map(makeImageBlock)
  setImageHolderInnards(ls.join(''))
}

function makeImageBlock(filename) {
  return `
  <div class="image">
    <p class="link"> ${url + '/images/view/' + filename} </p>
    <img src="${url}/images/view/${filename}" />
    <button onclick="deleteFile('${filename}')"> Delete </button>
  </div>`
}


function setImageHolderInnards(htmlData) {
  document.getElementById('imageHolder').innerHTML = htmlData
}

function deleteFile(filename) {
  if (confirm(`Are you sure you want to delete file ${filename}`)) {
    fetch(`${url}/images/delete/${filename}`, {
          method: 'DELETE'
      }).then(response => location.reload())
  }
}
