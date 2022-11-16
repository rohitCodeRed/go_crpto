// Page loads - script is called

var USERS=['Rahul','Coder',"Hacker"];
var CURRENT_USER_INFO={"name":"Rohit","url":"localhost:4000"};

var NODES=[
    {"address":"localhost:4001","name":"Rahul"},
    {"address":"localhost:4002","name":"Coder"},
    {"address":"localhost:4003","name":"Hacker"},
]
var TRANSACTIONS=[
    {"sender":"Rohit","reciever":"Coder","amount":1},
    {"sender":"Rohit","reciever":"Coder","amount":1},
    {"sender":"Rohit","reciever":"Coder","amount":1},
    {"sender":"Coder","reciever":"Hacker","amount":1},
    {"sender":"Rohit","reciever":"Coder","amount":1}
];

var BLOCKS ={
    "hashes":[
        {"index":0,"hash":"12998c017066eb0d2a70b94e6ed3192985855ce390f321bbdb832022888bd251"},
        {"index":1,"hash":"12998c017066eb0d2a70b94e6ed3192985855ce390f321bbdb832022888bd251"},
        {"index":2,"hash":"12998c017066eb0d2a70b94e6ed3192985855ce390f321bbdb832022888bd251"}
    ],
    "chain":[
        {
            "index":0,
            "dateTime":"",
            "proof":34557,
            "previous_hash":"9b96a1fe1d548cbbc960cc6a0286668fd74a763667b06366fb2324269fcabaa4",
            "transactions":[]
        },
        {
            "index":1,
            "dateTime":"",
            "proof":34557,
            "previous_hash":"9b96a1fe1d548cbbc960cc6a0286668fd74a763667b06366fb2324269fcabaa4",
            "transactions":[]
        },
        {
            "index":2,
            "dateTime":"",
            "proof":0,
            "previous_hash":"9b96a1fe1d548cbbc960cc6a0286668fd74a763667b06366fb2324269fcabaa4",
            "transactions":[{"sender":"Rohit","reciever":"Coder","amount":1},
            {"sender":"Coder","reciever":"Hacker","amount":1},
            {"sender":"Rohit","reciever":"Coder","amount":1}]
        }
    ]
    
}

var sendCoinModal = new bootstrap.Modal(document.getElementById('sendCoinModal'), {
    keyboard: false
});
  
var transactionModal = new bootstrap.Modal(document.getElementById('tranxModal'), {
    keyboard: false
});

function refreshUlist(){
    let uList = document.querySelector('.dropdown-menu');
    if (uList.hasChildNodes()) {
        uList.textContent="";
        uList.append(...listUsers(NODES));
    }else{
        uList.append(...listUsers(NODES));
    }

    var nodes= document.querySelectorAll('.dropdown-item');
    for(var i = 0 ; i < nodes.length ; i++) {
        nodes[i].addEventListener('click', function() {
            //alert('You clicked ' + this.getAttribute('data-reciever'));
            let pRName= document.getElementById("recieverName");
            pRName.textContent = this.getAttribute('data-reciever-name');
            pRName.setAttribute('data-address',this.getAttribute('data-reciever-url'));

            let pSName= document.getElementById("senderName");
            pSName.textContent = this.getAttribute('data-sender-name');
            pSName.setAttribute('data-address',this.getAttribute('data-sender-url'));
            
            sendCoinModal.show();

        });
    }
}

function listUsers(data){
    let fullHtml='';
    data.forEach(element => {
        fullHtml = fullHtml+`<li><a class="dropdown-item" data-sender-name="${CURRENT_USER_INFO.name}" data-sender-url="${CURRENT_USER_INFO.url}" data-reciever-url="${element.address}" data-reciever-name="${element.name}" href="#">${element.name}</a></li>`;
    });

    let domElem = htmlToElements(fullHtml);
    //console.log("dom Elem: ",domElem);
    return domElem;
}

function htmlToElements(html) {
    var template = document.createElement('template');
    template.innerHTML = html;
    return template.content.childNodes;
}

//---------transaction code---------

function refreshTransaction(){
    let uTrans = document.querySelector('.transaction-list');
    if (uTrans.hasChildNodes()) {
        uTrans.textContent="";
        uTrans.append(...addTransactions(TRANSACTIONS));
    }else{
        uTrans.append(...addTransactions(TRANSACTIONS));
    }
}
function addTransactions(data){
    let fullHtml='';
    for(let i=0;i<data.length;i++){
        let tranx = data[i];
        fullHtml = fullHtml+`<div class="card">
        <div class="card-body">
          <h6 class="card-subtitle mb-2 text-muted">Un Mined Transaction #${i + 1}</h6>
          <div class="row">
            <div class="col">
              <div class="tElem">
                <center>
                <div><strong>${tranx.sender.toUpperCase() == CURRENT_USER_INFO.name.toUpperCase()? "you":tranx.sender}</strong></div>
                <div class="text-muted">(sender)</div>
                  </div></center>
            </div>
            <div class="col-1">
             <svg width="16" height="16" fill="currentColor" class="bi bi-arrow-right" viewBox="0 0 16 16">
                <path fill-rule="evenodd" d="M1 8a.5.5 0 0 1 .5-.5h11.793l-3.147-3.146a.5.5 0 0 1 .708-.708l4 4a.5.5 0 0 1 0 .708l-4 4a.5.5 0 0 1-.708-.708L13.293 8.5H1.5A.5.5 0 0 1 1 8z"/>
             </svg>
            </div>
            <div class="col">
              <div class="tElem"><center>
                <div><strong>${tranx.reciever.toUpperCase() == CURRENT_USER_INFO.name.toUpperCase()? "you":tranx.reciever}</strong></div>
                <div class="text-muted">(reciever)</div></center>
              </div>
            </div>
             <div class="col">
              <div class="tElem"><center>
                <div><strong>${tranx.amount.toString()}</strong></div>
                <div class="text-muted">(amount)</div></center>
              </div>
            </div>
          </div>
        </div>
      </div>`;
    }
    // data.forEach(tranx => {
        
    // });

    let domElem = htmlToElements(fullHtml);
    //console.log("dom Elem: ",domElem);
    return domElem;
}

//--Block code----------
function refreshBlocks(){
    let uBlock = document.querySelector('.outer-block');
    if (uBlock.hasChildNodes()) {
        uBlock.textContent="";
        uBlock.append(...addBlocks(BLOCKS));
    }else{
        uBlock.append(...addBlocks(BLOCKS));
    }


    var nodes= document.querySelectorAll('.blockCard');
    for(var i = 0 ; i < nodes.length ; i++) {
        nodes[i].addEventListener('click', function() {
            let pIndex = parseInt(this.getAttribute('data-block-index'));
            let pId = parseInt(this.getAttribute('data-block-index')) + 1;
            //alert('You clicked ' + this.getAttribute('data-block-index'));
            let pLabel= document.getElementById("tranxModalLabel");
            pLabel.textContent = `Mined Transactions for Block #${pId.toString()}`;

            let mineTranx = document.getElementById("mine-transactions");

            mineTranx.textContent="";
            //let pIndex = parseInt(this.getAttribute('data-block-index'));
            mineTranx.append(...addTransactions(getBlockTransaction(pIndex,BLOCKS.chain)));

            transactionModal.show();

        });
    }
}

function getBlockTransaction(index,data){
    let pIndex = data.findIndex(function(elem){
        return elem.index == index;
    });

    return data[pIndex]["transactions"];
}

function addBlocks(data){
    let fullHtml='';
    let hashes = data.hashes;
    let mappedHashes =  new Map();
    hashes.forEach(elem => {
        mappedHashes.set(elem.index, elem.hash);
    });

    let chains = data.chain;
    chains.sort(function(a, b){return b.index - a.index});

    for(let i=0;i<chains.length;i++){
        let block = chains[i];
        //console.log(block);
        if(i != 0){
            fullHtml = fullHtml+`<div class="arrow-up"><svg width="32" height="32" fill="currentColor" class="bi bi-arrow-up" viewBox="0 0 16 16">
            <path fill-rule="evenodd" d="M8 15a.5.5 0 0 0 .5-.5V2.707l3.146 3.147a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 1 0 .708.708L7.5 2.707V14.5a.5.5 0 0 0 .5.5z"/>
          </svg></div>`;
        }

        fullHtml = fullHtml+`<div class="card blockCard" data-block-index="${block.index.toString()}">
        <div class="card-body">
        <h5 class="card-title text-muted">Block #${(block.index + 1).toString()}</h5>
        <div class="block-info">
            <div class="row">
        <label class="col-3 col-form-label"><strong>Hash:</strong></label>
        <div class="col-9">
            <input type="text" readonly class="form-control-plaintext" value="${mappedHashes.get(block.index)}">
        </div>
        </div>
            <div class="row">
        <label class="col-3 col-form-label"><strong>Nonce:</strong></label>
        <div class="col-9">
            <input type="number" readonly class="form-control-plaintext" value="${block.proof.toString()}">
        </div>
        </div>
            <div class="row">
            <label class="col-3 col-form-label"><strong>Prev hash:</strong></label>
        <div class="col-9">
            <input type="text" readonly class="form-control-plaintext" value="${block.previous_hash}">
        </div>
        </div>
            <div class="row">
        <label class="col-3 col-form-label"><strong>Transactions:</strong></label>
        <div class="col-9">
            <input type="number" readonly class="form-control-plaintext" value="${block.transactions.length.toString()}">
        </div>
        </div>
        </div>
        </div>
        </div>`;
  
    }

    let domElem = htmlToElements(fullHtml);
    //console.log("dom Elem: ",domElem);
    return domElem;
}


function showAlert() {
    alert("Ive Been Clicked!!");
}

//--------------------------ALL events code start-----------------

var mineBtn = document.querySelector('.mineBtn');
mineBtn.addEventListener('click', function () {
    refreshTransaction();
    refreshBlocks();
});

  //var sendBtn = document.getElementById('sendMenuBtn');
var sendMenuBtn = document.getElementById('sendMenuBtn')
sendMenuBtn.addEventListener('show.bs.dropdown', function () {
    // do something...
    refreshUlist();
});

var sendCoinBtn = document.getElementById('sendCoinBtn')
sendCoinBtn.addEventListener('click', function () {
    let senderUrl = document.getElementById('senderName').getAttribute('data-address');
    let senderName = document.getElementById('senderName').textContent;
    let recieverName = document.getElementById('recieverName').textContent;
    let recieverUrl = document.getElementById('recieverName').getAttribute('data-address');
    let amount = document.getElementById('sendCoinInput').value.toString();

    console.log(senderName,senderUrl,recieverName,recieverUrl,amount)
});


//----------------------------Events code end--------------------




//----------Socket code start-----------------------------------



//----------Socket code end------------------------------------
 
    


//-------------------------------------------------------------------------------
  //var sendBtn = document.querySelector('')
  
  //Go to the page, find the parent element of the buttons
  //var buttonDiv = document.querySelector('.parent');
  
  // Add a new button to the page whenever the "Add a new button below." button
  // is pressed.
  
  // We refence the variable defined earlier which found our special new button on the page and assigned an event listener to trigger specifically on click that calls an anonymous function
  // addNewButton.addEventListener('click', function () {
  //   // assigning a var to the creation of a new dom node
  //   var newButton = document.createElement('button');
  //   // adding a class to our new dom node
  //   newButton.className = 'button';
  //   // Adding text content to the new dom node
  //   newButton.textContent = "New click me button!";
  //   // Spitting said dom node on to page
  //   buttonDiv.appendChild(newButton);
  // });
  
  // Bind an event to all of the "Click me!" buttons that shows an alert.
  
  //Go to the page, find all buttons within the parent div
  //var allButtons = document.querySelectorAll('.parent .button');
  
  // For loop that iterates through the pre-defined variable that went out and found all the buttons on the page at page load
  // for(var i = 0; i < allButtons.length; i++){
  //   allButtons[i].addEventListener('click', showAlert);
  // }
  