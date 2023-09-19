class Page {
    constructor(config) {
      this.container = document.getElementById('root');
      this.renderFunctions = Object.create(this);
      this.renderFunctions.header = this.createHeader;
      this.renderFunctions.list = this.createList;
      this.renderFunctions.form = this.createForm;
      this.renderFunctions.panel = this.createPanel;
      this.renderFunctions.buttons = this.createButtons;
      this.config = config;
      this.init()
    }
    getRandom() {
      return Math.floor(Math.random() * 90000) + 10000;
    }
    createFakeData(renderConfig) {
      const fackDatatype = {
        string: 'Lorem ipsum dolor sit, amet consectetur',
        number: 15,
        date: '2023-01-08'
      };
      const fackData = [];
      const fackDataObj = {};
      renderConfig.columns.forEach((column, index) => {
        fackDataObj[column] = fackDatatype[renderConfig.type[index]];
      });
      for (let i = 0; i < 21; i++) {
        const copyData = Object.assign({}, fackDataObj);
        fackData.push(copyData);
      }
      console.log(fackData);
      return fackData;
    }
    createRouter() {
      const thisObj = this;
      const route = new Router();
      const config = this.config;
      route.setRoute('', function () {
        console.log('首頁');
      });
      for (let pageName in config) {
        route.setRoute(`/${pageName}`, function () {
          console.log(pageName);
          const pageConfig = config[pageName];
          thisObj.renderPage(pageConfig);
        });
      }
      route.init();
    }
    createHeader(renderConfig, tabContent) {
      const headerBox = document.createElement('header');
      const header = document.createElement('h1');
      header.innerHTML = renderConfig.name;
      if (Array.isArray(renderConfig.class)) {
        header.classList.add(...renderConfig.class);
      } else {
        renderConfig.class && header.classList.add(renderConfig.class);
      }
      headerBox.classList.add(
        'd-flex',
        'justify-content-between',
        'align-items-center'
      );
      headerBox.appendChild(header);
      if (renderConfig.buttons) {
        this.createButtons(renderConfig.buttons, headerBox);
      }
      if (tabContent) {
        tabContent.appendChild(headerBox);
      } else {
        headerBox.classList.add('page-header');
        this.container.appendChild(headerBox);
      }
    }
    createMainBox(tabContent) {
      if (tabContent) return tabContent;
      let mainBox = this.container.querySelector('main');
      if (!mainBox) {
        mainBox = document.createElement('main');
        mainBox.classList.add('card', 'card-body', 'content');
        this.container.appendChild(mainBox);
      }
      return mainBox;
    }
    createList(renderConfig, tabContent) {
      const mainBox = this.createMainBox(tabContent);
      const list = document.createElement('table');
      list.id = `data-table-${this.getRandom()}`;
      list.classList.add('table', 'table-striped', 'table-bordered');
      mainBox.appendChild(list);
      const columns = renderConfig.columns.map((column) => {
        return {
          title: column,
          data: column
        };
      });
      const fackData = this.createFakeData(renderConfig);
      const table = new DataTable(`#${list.id}`, {
        autoWidth: true,
        scrollY: '56vh',
        scrollResize: true,
        scrollCollapse: true,
        columns,
        data: fackData
      });
    }
    createPanel(renderConfig) {
      this.createTabGrid(renderConfig);
    }
    createTabItem(targetId, name, isActive) {
      const tabItem = document.createElement('li');
      const tabBtn = document.createElement('span');
      tabItem.classList.add('nav-item');
      isActive && tabBtn.classList.add('active');
      tabBtn.classList.add('nav-link');
      tabBtn.dataset.toggle = 'tab';
      tabBtn.dataset.target = `#tab-content-${targetId}`;
      tabBtn.setAttribute('type', 'button');
      tabBtn.innerHTML = name || 'no name tab';
      tabItem.appendChild(tabBtn);
      return tabItem;
    }
    createTabContent(id, isActive) {
      const tabContent = document.createElement('div');
      tabContent.classList.add('tab-pane', 'fade');
      isActive && tabContent.classList.add('show', 'active');
      tabContent.id = `tab-content-${id}`;
      return tabContent;
    }
    createTabGrid(panelConfig) {
      const mainBox = this.createMainBox();
      const thisObj = this;
      const tabGrid = document.createElement('div');
      const tabList = document.createElement('ul');
      const tabContentContainer = document.createElement('div');
      tabGrid.appendChild(tabList);
      tabGrid.appendChild(tabContentContainer);
      mainBox.appendChild(tabGrid);
      panelConfig.forEach((config, index) => {
        const tabId = thisObj.getRandom();
        const tabItem = thisObj.createTabItem(tabId, config.name, !index);
        const tabContent = thisObj.createTabContent(tabId, !index);
        tabList.appendChild(tabItem);
        tabContentContainer.appendChild(tabContent);
        if (config.content) {
          thisObj.renderPage(config.content, tabContent);
        }
      });
      tabGrid.classList.add('tab-grid');
      tabList.classList.add('nav', 'nav-tabs');
      tabGrid.style.marginTop = '1rem';
      tabContentContainer.classList.add('tab-content');
      return tabGrid;
    }
    createColumn(columnConfig) {
      const thisObj = this;
      const name = columnConfig[0];
      const type = columnConfig[1];
      const width = columnConfig[2];
      const source = columnConfig[3];
      const columnId = `${type}-${this.getRandom()}`;
      const formGroup = document.createElement('div');
      const label = document.createElement('label');
      label.innerHTML = name;
      label.setAttribute('for', columnId);
      formGroup.classList.add('form-group', 'p-1');
      formGroup.appendChild(label);
      let column = null;
      switch (type) {
        case 'input':
          column = document.createElement('input');
          break;
        case 'number':
          column = document.createElement('input');
          column.type = 'number';
          break;
        case 'date':
          column = document.createElement('input');
          column.type = 'date';
          break;
        case 'textarea':
          column = document.createElement('textarea');
          break;
        case 'select':
          column = document.createElement('select');
          source.forEach((data) => {
            const option = new Option(data.text, data.value);
            column.appendChild(option);
          });
          break;
        case 'radio':
        case 'checkbox':
          label.removeAttribute('for', columnId);
          label.classList.add('d-block');
          column = [];
          source.forEach((data) => {
            const formCheckGroup = document.createElement('div');
            formCheckGroup.classList.add('form-check', 'form-check-inline');
            const radioId = `${name}-${thisObj.getRandom()}`;
            const label = document.createElement('label');
            const radio = document.createElement('input');
            radio.name = name;
            radio.id = radioId;
            radio.value = data.value;
            radio.type = type;
            radio.classList.add('form-check-input');
            label.innerHTML = data.text;
            label.setAttribute('for', radioId);
            label.classList.add('form-check-lebel');
            formCheckGroup.appendChild(radio);
            formCheckGroup.appendChild(label);
            column.push(formCheckGroup);
          });
          break;
      }
      if (Array.isArray(column)) {
        column.forEach((item) => {
          formGroup.appendChild(item);
        });
      } else {
        column.id = columnId;
        column.classList.add('form-control');
        formGroup.appendChild(column);
      }
      if (width) {
        formGroup.style.width = isNaN(+width) ? width : `${width}px`;
      }
      return formGroup;
    }
    createForm(renderConfig, tabContent) {
      const thisObj = this;
      const mainBox = this.createMainBox(tabContent);
      const form = document.createElement('form');
      form.classList.add('flex-form');
      const columns = renderConfig.columns;
      columns.forEach((columnConfig) => {
        const column = thisObj.createColumn(columnConfig);
        form.appendChild(column);
      });
      this.createButtons(
        [
          {
            name: 'Submit',
            class: ['btn', 'btn-success'],
            event: function () {
              console.log('Submit')
              form.submit();
            }
          }
        ],
        form
      );
      mainBox.appendChild(form);
    }
    renderPage(pageConfig, tabContent = null) {
      if (!tabContent) {
        this.container.innerHTML = '';
      }
      for (let renderName in pageConfig) {
        const renderConfig = pageConfig[renderName];
        this.renderFunctions[renderName](renderConfig, tabContent);
      }
    }
    createButtons(renderConfig, target) {
      const buttonBox = document.createElement('div');
      renderConfig.forEach((buttonConfig) => {
        const button = document.createElement('button');
        button.innerHTML = buttonConfig.name || 'button';
        if (Array.isArray(buttonConfig.class)) {
          button.classList.add(...buttonConfig.class);
        }
        button.addEventListener('click', buttonConfig.event);
        buttonBox.appendChild(button);
      });
      target.appendChild(buttonBox);
    }
    init() {
      this.createRouter();
      return;
      const root = document.getElementById('root');
      if (!root) return;
      if (this.config.header) {
        const header = createHeader(this.config.header);
        root.appendChild(header);
      }
    }
  }
  