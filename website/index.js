let allServices = [];
let allInstances = [];
let lastSelectedService = '';

function fillServiceTable() {
    // Clear table
    tableBody.innerHTML = '';
    // Fill table
    for (const service of allServices) {
        /*
        <tr class="table-success">
            <th scope="row">1</th>
            <td>Mark</td>
            <td>Otto</td>
            <td>@mdo</td>
        </tr>
        */
        let tr = document.createElement('tr');
        // Count all instances that have this serviceId
        let amountInstances = 0;
        allInstances.forEach((e) => {
            if(e.serviceid === service.id) {
                amountInstances++;
            }
        });
        if(amountInstances === 0) {
            tr.className = 'table-danger';
        }else {
            tr.className = 'table-success';
        }

        let th1 = document.createElement('th');
        th1.setAttribute('scope', 'row');
        // Make serviceId clickable to ope instances
        let a = document.createElement('a');
        a.className = 'link-primary'; a.innerText = service.id;
        a.addEventListener('click', () => fillInstancesTable(service.id));
        th1.appendChild(a);

        let td2 = document.createElement('td');
        td2.innerText = service.serviceType;

        let td3 = document.createElement('td');
        if(amountInstances === 0) {
            td3.innerText = 'DOWN';
        }else {
            td3.innerText = 'UP';
        }

        let td4 = document.createElement('td');
        td4.innerText = amountInstances;

        tr.appendChild(th1);
        tr.appendChild(td2);
        tr.appendChild(td3);
        tr.appendChild(td4);

        tableBody.appendChild(tr);

    }
}

function fillInstancesTable(serviceId) {
    lastSelectedService = serviceId;

    // Set section header
    document.getElementById('currentService').innerText = 'Service: ' + serviceId;

    // Clear table
    let tableBodyInstances = document.getElementById('tableBodyInstances');
    tableBodyInstances.innerHTML = '';

    // Current time in unix seconds
    let currentTime = Math.floor(Date.now() / 1000);

    // Find all instances of serviceId
    let instances = allInstances.filter(e => e.serviceid === serviceId)
    console.log(serviceId);
    console.log(instances);

    if(instances.length !== 0) {
        for(let i = 0; i < instances.length; i++) {
            let tr = document.createElement('tr');
            if(currentTime - instances[i].lastHeartbeat > 90) {
                tr.className = 'table-danger';
            } else if(currentTime - instances[i].lastHeartbeat > 30) {
                tr.className = 'table-warning';
            }else {
                tr.className = 'table-success';
            }

            let th1 = document.createElement('th');
            th1.setAttribute('scope', 'row');
            th1.innerText = instances[i].id;

            let td2 = document.createElement('td');
            td2.innerText = instances[i].ip;

            let td3 = document.createElement('td');
            td3.innerText = instances[i].port;

            let td4 = document.createElement('td');
            //td4.innerText = value.map(e => JSON.stringify(e)).join('\n') 
            td4.innerText = (currentTime - instances[i].lastHeartbeat) + ' sec ago';

            tr.appendChild(th1);
            tr.appendChild(td2);
            tr.appendChild(td3);
            tr.appendChild(td4);

            tableBodyInstances.appendChild(tr);

        }
    }
}

function getInstances() {
    fetch('http://localhost:2000/instances')
        .then(response => response.json())
        .then(data => {
            allInstances = data
    });
}

function getServices() {
    // Get all serviceIds
    fetch('http://localhost:2000/apps')
        .then(response => response.json())
        .then(data => {
            allServices = data;
    });
}

function update() {
    getServices();
    getInstances();
    fillServiceTable();
    fillInstancesTable(lastSelectedService);
}

let tableBody = document.getElementById('tableBody');
update()
fillServiceTable()

setInterval(update, 2000);