() => {
    const config = { attributes: true, childList: true };
    const callback = function (mutationList, observer) {
        // console.log(mutationList);

        if (mutationList[0]?.addedNodes.length > 0) {
            mutationList[0].addedNodes[0].children[0].click();
        }
    }

    const listCallBack = function (mutationList) {
        // console.log(mutationList);
        if (mutationList.length != 20) {
            mutationList.forEach(x => {
                // console.log(x.addedNodes);
                if (x.addedNodes.length > 0) {
                    if (x.addedNodes[0].nodeName == "LI") {
                        var project = {};
                        var list = x.addedNodes[0].children[0];
                        // console.log(list);
                        // debugger;
                        project.link = list.href;
                        project.skills = list.children[0].querySelector('.info-card-skillsContainer').children[1].innerText;
                        project.title = list.children[0].querySelector('.info-card-title').innerText;
                        project.description = list.children[0].querySelector('.info-card-description').innerText;
                        project.biddingPrice = list.children[0].querySelector('.info-card-price').children[0].innerText;
                        project.currency = list.children[0].querySelector('.info-card-price-type').children[0].innerText;
                        project.selected = false;

                        console.log(project,list)

                        window.GetProject(project);
                    }
                }
            })
        }
    }

    const listNode = document.getElementsByClassName('search-result-list')[0];
    const targetNode = document.getElementById('search-results');

    const listObserver = new MutationObserver(listCallBack);
    const observer = new MutationObserver(callback);

    if(!listNode) {
        // window.requestAnimationFrame(() => {
        //     listNode = document.getElementsByClassName('search-result-list')[0];
        // })
     setTimeout(() => {
        listNode = document.getElementsByClassName('search-result-list')[0];
        targetNode = document.getElementById('search-results');
    });   
    } else {
        listObserver.observe(listNode, config);
        observer.observe(targetNode, config);
    }
}