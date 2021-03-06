import BaseView from '../BaseView.js';
import testTemplate from './TestView.hbs';
import Input from '../Common/Input/Input';
import testStyles from './TestView.scss';
import textStyles from './../Common/TextArea/TextArea.scss';
import imgStyles from './../Common/Img/Img.scss';
import buttonStyles from './../Common/Button/Button.scss';
import linkStyles from './../Common/Link/Link.scss';



class Project {
    constructor({name = ''} = {}) {
        this.name = name;
    }
}

class TicketType {
    constructor({name = ''} = {}) {
        this.name = name;
    }
}

class User {
    constructor({name = ''} = {}) {
        this.name = name;
    }
}

class Ticket {
    constructor({number=-1, name='', project='', type='', linkedNum=-1, author='', executor='', description='', complexity=-1, attachments=''}) {
        this.number = number;
        this.name = name;
        this.project = project;
        this.type = type;
        this.linkedNum = linkedNum;
        this.author = author;
        this.executor = executor;
        this.description = description;
        this.complexity = complexity;
        this.attachments = attachments;
    }
}

/**
 * @class TestView
 * @extends BaseView
 * @classdesc Class for showing test
 */
class TestView extends BaseView {
    show = () => {
        this.backlog = [];
        this.render();
    }

    render = () => {
        this.users = [
            new User({
                name: 'Employee_A',
            }),
            new User({
                name: 'Employee_B',
            }),
            new User({
                name: 'Employee_C',
            }),
        ];
        this.projects = [
            new Project({
                name: 'Project_A',
            }),
            new Project({
                name: 'Project_B',
            }),
            new Project({
                name: 'Project_C',
            }),
        ];
        this.ticketTypes = [
            new TicketType({
                name: 'task',
            }),
            new TicketType({
                name: 'story',
            }),
            new TicketType({
                name: 'bug',
            }),
        ];
        this.parent.innerHTML = '';
        const template = testTemplate({
            projects: this.projects,
            types: this.ticketTypes,
            authors: this.users,
            executors: this.users,
            numbers: this.backlog.length,
            styles: testStyles,
            textStyles: textStyles,
            buttonStyles: buttonStyles,
            imgStyles: imgStyles,
            linkStyles: linkStyles,
        });
        this.cache = new DOMParser().parseFromString(template, 'text/html').getElementById('test-block');
        const form = this.cache.getElementsByClassName('form')[0];
        const backlog = this.backlog;
        const view = this;
        form.addEventListener('submit', (evt) => {
            evt.preventDefault();
            const formData = new FormData(form);
            console.log()
            backlog.push(
                new Ticket({
                    name: formData.get('name'),
                    project: formData.get('project'),
                    type: formData.get('type'),
                    linkedNum: formData.get('linkedNum'),
                    author: formData.get('author'),
                    executor: formData.get('executor'),
                    description: formData.get('description'),
                    complexity: formData.get('complexity'),
                    attachments: document.getElementById('file').value,
                    number: view.backlog.length,
                }),
            );
            view.render();
        });
        this.parent.appendChild(this.cache);
        document.getElementById('tickets').innerHTML = JSON.stringify(this.backlog, null, 0);
    };
}

export default TestView;
