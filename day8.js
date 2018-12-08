const input = require('./input')[0].split(' ').map(v => parseInt(v))

function parseNode(data,offset=0) {
  const childCount = data[offset]
  const metadataCount = data[offset+1]
  const children = []
  let off = 2
  for(let i = 0; i < childCount; i++ ){
    const node = parseNode(data, offset + off)
    children.push(node)
    off += node.size
  }
  const size = off + metadataCount
  const metadata = data.slice(offset + off, offset + off + metadataCount)
  const sum = metadata.reduce((a,b) => a + b, 0) + children.reduce((a,b) => a + b.sum, 0)
  const value = childCount ? metadata.reduce((a,b) => a + (children[b-1] ? children[b-1].value : 0), 0) : sum
  return { children, metadata, size, sum, value }
}

const tree = parseNode(input)
//console.log(JSON.stringify(tree))
console.log(tree.sum, tree.value)