() => {
    let projectList = document.getElementsByClassName('search-result-list')[0].children;

                let projects = [];

                for (let i = 0; i < projectList.length - 1; i++) {

                    let element = projectList[i].children[0];

                    let project = {};
                    
                    

                    let skillChildren = element.children[0].querySelector('.info-card-skills').children;

                    let skills = (() => {let skillsArr = []; for(let skill of skillChildren) skillsArr.push(skill.innerText); return skillsArr})().join(",")


                    project.link = element.href;
                    project.skills = skills;
                    project.title = element.children[0].querySelector('.info-card-title').innerText;
                    project.description = element.children[0].querySelector('.info-card-description').innerText;
                    project.biddingPrice = element.children[0].querySelector('.info-card-price').children[0].innerText;
                    project.currency = element.children[0].querySelector('.info-card-price-type').children[0].innerText;
                    project.selected = false;

                    projects.push(project)
                }

                window.GetProjectList(projects);
}