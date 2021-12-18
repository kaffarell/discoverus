let services = new Map();
let lastSelectedService = '';

function fillServiceTable() {
    // Clear table
    tableBody.innerHTML = '';
    // Fill table
    for (const [key, value] of services.entries()) {
        /*
        <tr class="table-success">
            <th scope="row">1</th>
            <td>Mark</td>
            <td>Otto</td>
            <td>@mdo</td>
        </tr>
        */
        let tr = document.createElement('tr');
        if(value.length === 0) {
            tr.className = 'table-danger';
        }else {
            tr.className = 'table-success';
        }

        let th1 = document.createElement('th');
        th1.setAttribute('scope', 'row');
        // Make serviceId clickable to ope instances
        let a = document.createElement('a');
        a.className = 'link-primary'; a.innerText = key;
        a.addEventListener('click', () => fillInstancesTable(key));
        th1.appendChild(a);

        let td2 = document.createElement('td');
        td2.innerText = 'service';

        let td3 = document.createElement('td');
        if(value.length === 0) {
            td3.innerText = 'DOWN';
        }else {
            td3.innerText = 'UP';
        }

        let td4 = document.createElement('td');
        //td4.innerText = value.map(e => JSON.stringify(e)).join('\n') 
        td4.innerText = value.length;

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

    let instances = services.get(serviceId);
    // Current time in unix seconds
    let currentTime = Math.floor(Date.now() / 1000);

    if(instances) {
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

function getInstances(serviceId) {

    fetch('http://localhost:2000/apps/' + serviceId)
        .then(response => response.json())
        .then(data => {
            services.set(serviceId, data)
    });
}

function getServices() {
    // Get all serviceIds
    let serviceIdArray = [];
    fetch('http://localhost:2000/apps')
        .then(response => response.json())
        .then(data => {
            serviceIdArray = data;
            for(let i = 0; i < serviceIdArray.length; i++) {
                getInstances(serviceIdArray[i]);
            }
    });
}

function update() {
    getServices();
    fillServiceTable();
    fillInstancesTable(lastSelectedService);
}

let tableBody = document.getElementById('tableBody');
update()
fillServiceTable()

setInterval(update, 2000);