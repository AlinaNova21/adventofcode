canvas.width = window.innerWidth
canvas.height = window.innerHeight
slider.max = 100000

class Graph {
  constructor() {
    this.edges = []
    this.nodes = {}
    this.links = []
  }
  addEdge(a, b, opts) {
    if(this.nodes[a] && this.nodes[b]) {
      this.edges.push([a,b,opts])
      this.edges.push([b,a,opts])
      this.links.push({ source: a, target: b, weight: opts.weight })
      this.nodes[a].edges[b] = opts
      this.nodes[b].edges[a] = opts
      this.nodes[a].neighbors.push(b)
      this.nodes[b].neighbors.push(a)
    }
  }
  addNode(node) {
    node.edges = {}
    node.neighbors = []
    this.nodes[node.id] = node
  }
  neighbors(node) {
    return this.nodes[node].neighbors
    const ret = []
    return this.edges
      .filter(e => e[0] === node)
      .map(e => e[1])
  }
  cost(a,b) {
    return this.nodes[a].edges[b].weight
    const edge = this.edges.find(e => e[0] === a && e[1] === b)
    return edge[2].weight * 1
  }
  heuristic(a, b) {
    // return 0
    const nA = this.nodes[a]
    const nB = this.nodes[b]
    const dist = Math.abs(nA.pos[0] - nB.pos[0]) + Math.abs(nA.pos[1] - nB.pos[1]) + Math.abs(nA.pos[2] - nB.pos[2])
    return dist
  }
}

function* pathfind(graph, start, goal) {
  console.log(graph)
  // return { cost: 0, path: [] }
  const frontier = new buckets.PriorityQueue((a,b) => b[1] - a[1])
  frontier.enqueue([start, 0])
  const cameFrom = {}
  const costSoFar = {}
  cameFrom[start] = null
  costSoFar[start] = 0
  while(!frontier.isEmpty()) {
    const [current] = frontier.dequeue()
    // if (current === goal) break // Early aborting doesn't guarantee shortest path
      
    for(const next of graph.neighbors(current)) {
      const newCost = costSoFar[current] + graph.cost(current, next)
      if (!costSoFar[next] || newCost < costSoFar[next]) {
        costSoFar[next] = newCost
        const priority = newCost + graph.heuristic(goal, next)
        frontier.enqueue([next, priority])
        cameFrom[next] = current
      }
    }
    yield { current, start, goal, cameFrom, costSoFar, frontier }
  }
  const path = []
  let current = goal
  while(current !== start) {
    path.push(current)
    current = cameFrom[current]
  }
  path.push(start)
  path.reverse()
  return { cost: costSoFar[goal], path: path.map(id => graph.nodes[id]) }
}


const clay = []
let world
async function run() {
  const res = await fetch('input/day22')
  const data = await res.text()
  const input = data.split('\n').filter(Boolean)
  const SCALE = 10
  const [[depth],target] = input.map(v => v.split(' ')[1].split(',').map(v => +v))
  // const [depth,target] = [510,[10,10]]
  const w = target[0] + 20
  const h = target[1] + 20

  const grid = new Array(w*h)
  grid.fill(0)

  function calcGeoIndex(ind) {
    const [x,y] = toXY(ind)
    if(x === 0 && y === 0) return 0
    if(x === target[0] && y === target[1]) return 0
    if(y === 0) return x * 16807
    if(x === 0) return y * 48271
    const [aInd,bInd] = [toInd([x-1,y]), toInd([x,y-1])]
    const a = grid[aInd] || calcErosionLevel(aInd)
    const b = grid[bInd] || calcErosionLevel(bInd)
    return a * b
  }

  function calcErosionLevel(ind) {
    return (calcGeoIndex(ind) + depth) % 20183
  }

  grid.forEach((v,i) => {
    grid[i] = calcErosionLevel(i)
  })

  const smallestRect = target
  let smallest = 0
  for(let y = 0; y <= target[1]; y++) {
    for(let x = 0; x <= target[0]; x++) {
      smallest += grid[toInd([x,y])] % 3
    }
  }

  const graph = new Graph()
  // 0        1       2
  // neither  torch   climb
  // rocky    wet     narrow
  // climb    climb   torch
  // torch    neither neither

  for(let y = 0; y < h; y++) {
    for(let x = 0; x < w; x++) {
      const v = grid[toInd([x,y])] % 3
      for(let i = 0; i < 3; i++) {
        if(i === v) continue
        graph.addNode({ id: `${x},${y}-${i}`, i, v, x, y, z: i, pos: [x,y,i] })
        if(x) {
          const v2 = grid[toInd([x-1,y])] % 3
          if(i !== v2) {
            graph.addEdge(`${x},${y}-${i}`, `${x-1},${y}-${i}`, { weight: 1 })
          }
        }
        if(y) {
          const v2 = grid[toInd([x,y-1])] % 3
          if(i !== v2) {
            graph.addEdge(`${x},${y}-${i}`, `${x},${y-1}-${i}`, { weight: 1 })
          }
        }
      }
      if(v === 0) graph.addEdge(`${x},${y}-1`, `${x},${y}-2`, { weight: 7 })
      if(v === 1) graph.addEdge(`${x},${y}-0`, `${x},${y}-2`, { weight: 7 })
      if(v === 2) graph.addEdge(`${x},${y}-0`, `${x},${y}-1`, { weight: 7 })
    }
  }

  const customBase = document.createElement('custom');
  const custom = d3.select(customBase)
  function databind() {
    const nodeArr = Object.values(graph.nodes)
    const nodes = nodeArr.map(n => ({ id: n.id }))
    const links = graph.links
    const sim = d3.forceSimulation()
      .force('link', d3.forceLink().id(d => d.id).strength(0.1).distance(d => d.weight))
      .force('charge', d3.forceManyBody().strength(-5))
      .force('center', d3.forceCenter(canvas.width/2,canvas.height/2))
      .force('collision', d3.forceCollide().radius(d => 10))
    sim.nodes(nodes)
    sim.force('link').links(links)
    
    return { sim, nodes, links }
  }
  // const force = databind()
  const pathfinder = pathfind(graph, '0,0-1', `${target[0]},${target[1]}-1`)
  let path = []
  function pathStep() {
    for(let i = 0; i < 100; i++) {
      const { value: { cameFrom, costSoFar, start, goal, path: finalPath, cost } = {}, value, done } = pathfinder.next()
      if(finalPath) {
        path = finalPath.map(n => n.pos)
        path.cost = cost
      }
      if(done) break
      let current = value.current
      path = []
      while(current !== start) {
        path.push(current)
        current = cameFrom[current]
      }
      path.push(start)
      path.reverse()
      path = path.map(id => graph.nodes[id].pos)
      path.cost = costSoFar[value.current]
    }
  }
  function render() {
    const targetY = Math.max(+slider.value, (path.slice(-1)[0][1] * SCALE) - (window.innerHeight / 2))
    {
      const v = +slider.value
      const dist = targetY - v
      slider.value = v + (dist / 20)
    }
    const ctx = canvas.getContext('2d')
    ctx.clearRect(0,0,canvas.width,canvas.height)
    ctx.save()
    ctx.font = '10pt mono'
    ctx.translate(0, -slider.value)
    // ctx.translate(0,100)
    const colors = [
      'gray',
      'green',
      'blue',
      '#ccc',
      'orange',
      'black'
    ]
    const gridQueue = [[],[],[]]
    const startY = Math.max(0, Math.floor(+slider.value / SCALE))
    const endY = Math.min(h, startY + Math.ceil(canvas.height / SCALE))
    for(let y = startY; y < endY; y++) {
      for(let x = 0; x < w; x++) {
        const ind = toInd([x,y])
        const v = grid[ind] % 3
        gridQueue[v].push([x,y])
      }
    }
    for(const ind in gridQueue) {
      const queue = gridQueue[ind]
      ctx.fillStyle = colors[ind]
      ctx.beginPath()
      for(const [x,y] of queue) {
        ctx.rect(x*SCALE,y*SCALE,SCALE,SCALE)
      }
      ctx.fill()
      ctx.fillStyle = 'white'
      for(const [x,y] of queue) {
        // ctx.fillText(ind,x*SCALE,y*SCALE + 10)
      }
    }
    ctx.save()
    const toolColors = [
      'black',
      'yellow',
      'silver',
    ]
    let tool = 1
    let move = 0
    let toolSwap = 0
    ctx.strokeStyle = toolColors[tool]
    ctx.lineWidth = 2 //SCALE / 4
    ctx.beginPath()
    ctx.translate(SCALE/2, SCALE/2)
    ctx.moveTo(0,0)
    let last = [-1,-1]
    path.forEach(([x,y,i]) => {
      // if(last[0] === x && last[1] === y) return
      last = [x,y]
      ctx.lineTo(x * SCALE, y * SCALE)
      ctx.fillStyle = 'black'
      // ctx.fillText(i, x * SCALE, y * SCALE)
    })
    ctx.stroke()
    ctx.restore()
    ctx.restore()
    ctx.save()
    ctx.fillStyle = 'white'
    ctx.font = '16pt Mono'
    ctx.translate(400,20)
    ctx.fillText(`Smallest: ${smallest}`, 0, 0)
    const min = (toolSwap * 7) + move
    // ctx.fillText(`Minutes: ${min}`, 0, 20)
    // ctx.fillText(`ToolSwap: ${toolSwap}`, 0, 40)
    // ctx.fillText(`Moves: ${move}`, 0, 60)
    ctx.fillText(`Cost: ${path.cost}`, 0, 20)
    ctx.restore()
    ctx.save()
    /** /
    ctx.strokeStyle = 'gray'
    ctx.beginPath()
    force.links.forEach(link => {
      ctx.moveTo(link.source.x, link.source.y)
      ctx.lineTo(link.target.x, link.target.y)
    })
    ctx.stroke()

    ctx.fillStyle = 'white'
    ctx.beginPath()
    force.nodes.forEach(node => {
      ctx.moveTo(node.x,node.y)
      ctx.arc(node.x, node.y, 5, 0, Math.PI*2)
    })
    ctx.fill()
    ctx.fillStyle = 'red'
    force.nodes.forEach(node => {
      ctx.fillText(node.id,node.x,node.y)
    })
    /**/
    ctx.restore()
  }
  const scene = new THREE.Scene()
  const camera = new THREE.PerspectiveCamera(75, window.innerWidth/window.innerHeight, 0.1, 1000)
  const renderer = new THREE.WebGLRenderer({ alpha: true })
  renderer.setSize(window.innerWidth, window.innerHeight)
	document.body.appendChild(renderer.domElement)

  const boxGeometry = new THREE.BoxGeometry(0.4, 0.4, 0.4)
  const sphereGeometry = new THREE.SphereGeometry(0.3)
  const layerMaterial = [
    0xff0000,
    0x00ff00,
    0x0000ff,
  ].map(color => new THREE.MeshLambertMaterial({ color, transparent: true, opacity: 0.4 }))
  const lineMaterial = new THREE.LineBasicMaterial({ color: 0x0000ff, transparent: true, opacity: 0.5 })
  const pathMaterial = new THREE.LineBasicMaterial({ color: 0xffff00 })

  const light = new THREE.AmbientLight(0x404040,0.5) // soft white light
  scene.add(light)

  const directionalLight = new THREE.DirectionalLight(0xffffff, 0.8)
  directionalLight.position.set(-10,-10,50)
  // scene.add(directionalLight)
  
  const pointLight = new THREE.PointLight(0xffffff, 5, 50, 2)
  pointLight.position.set(0, 0, 5)
  scene.add(pointLight)
  // const pointLightHelper = new THREE.PointLightHelper(pointLight, 5)
  // scene.add(pointLightHelper)

  const layerGeometry = [
    new THREE.Geometry(),
    new THREE.Geometry(),
    new THREE.Geometry()
  ]

  for(const node of Object.values(graph.nodes)) {
    const mesh = new THREE.Mesh(boxGeometry, layerMaterial[node.z])
    mesh.position.x = node.x
    mesh.position.y = node.y
    mesh.position.z = node.z * 7
    mesh.updateMatrix()
    layerGeometry[node.z].mergeMesh(mesh)
  }
  // nodeGeometry.mergeVertices()
  scene.add(new THREE.Mesh(layerGeometry[0], layerMaterial[0]))
  scene.add(new THREE.Mesh(layerGeometry[1], layerMaterial[1]))
  scene.add(new THREE.Mesh(layerGeometry[2], layerMaterial[2]))
    
  for(const link of graph.links) {
    const a = graph.nodes[link.source]
    const b = graph.nodes[link.target]
    const geo = new THREE.Geometry()
    geo.vertices.push(
      new THREE.Vector3(a.x, a.y, a.z * 7),
      new THREE.Vector3(b.x, b.y, b.z * 7)
    )
    const line = new THREE.Line(geo, lineMaterial)
    // scene.add(line)
  }

  const pathGeometry = new THREE.BufferGeometry()
  const pathPositions = new Float32Array(2000 * 3)
  pathGeometry.addAttribute('position', new THREE.BufferAttribute(pathPositions, 3))
  pathGeometry.setDrawRange(0, 0)
  
  const pathLine = new THREE.Line(pathGeometry, pathMaterial)
  scene.add(pathLine)
  
  camera.position.z = 30
  camera.position.x = 10
  camera.position.y = 0
  camera.up = new THREE.Vector3(0,0,1);
  camera.lookAt(new THREE.Vector3(10,50,7))
  // camera.rotation.y = 1
  // camera.rotation.z = 1

  function threeRender() {
    const [[x,y,z]] = path.slice(-1)
    const diff = Math.max(0, ((y - 20) - camera.position.y) / 50)
    camera.position.y += diff
    if(diff === 0 && (y + 200) < h) {
      camera.position.y += 0.1
    }
    // pointLight.position.set(camera.position)
    pointLight.position.x = 10
    pointLight.position.z = 20
    pointLight.position.y = camera.position.y
    renderer.render(scene, camera)
  }
  renderLoop()
  // btn.addEventListener('click', function () {
  // world.step()
    // renderLoop()
  // })
  // btn.innerHTML = 'start'
  function renderLoop() {
    requestAnimationFrame(renderLoop)
    pathStep()
    // console.log(path,pathPositions)
    const positions = pathLine.geometry.attributes.position.array
    path.forEach(([x,y,z],i) => {
      positions[(i*3) + 0] = x
      positions[(i*3) + 1] = y
      positions[(i*3) + 2] = z * 7
    })
    pathLine.geometry.setDrawRange(0, path.length)
    pathLine.geometry.attributes.position.needsUpdate = true // required after the first render
    
    pathPositions.needUpdate = true
    pathGeometry.verticesNeedUpdate = true
    pathLine.frustumCulled = false
    render()
    threeRender()
  }
  
  function toXY (ind) {
    const rx = ind % w
    const ry = Math.floor(ind / w)
    return [rx, ry]
  }
  
  function toInd([x,y]) {
    return (y * w) + x
  }
}

run().catch(err => console.error(err))
