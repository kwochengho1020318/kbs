class Core {
  constructor(config) {
    this.randomList = [];
  }
  createEl(tagName = 'div', content = '', attrs = {}, children = []) {
    const el = document.createElement(tagName);
    el.textContent = content;
    for (const attrName in attrs) {
      el.setAttribute(attrName, attrs[attrName]);
    }
    children.forEach((child) => {
      el.append(child);
    });
    return el;
  }
  getDom(target = null, type = '') {
    type && (type = type.toUpperCase());
    let element = null;
    if (!target) {
      return console.warn('getDom: parameter(target) is not defined');
    }
    if (target instanceof HTMLElement) {
      element = target;
    } else if (typeof target === 'string') {
      element = document.getElementById(target);
    }
    if (!element || (element && type && element.tagName !== type)) {
      console.warn('getDom: parameter(target) is not correct');
      element = null;
    }
    return element;
  }
  async getData(apiUrl = '', query = {}) {
    const api = this.getQueryUrl(apiUrl, query);
    const response = await fetch(api);
    if (response.ok) {
      const responseData = await response.json();
      return responseData;
    } else {
      alert('Failed to get data, please try again later.');
      return null;
    }
  }
  async postData(apiUrl = '', para = {}) {
    const response = await fetch(apiUrl, {
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(para)
    });
    if (response.ok) {
      const responseData = await response.json();
      return responseData;
    } else {
      alert('Failed to get data, please try again later.');
      return null;
    }
  }
  async putData(apiUrl = '', para = {}) {
    const response = await fetch(apiUrl, {
      method: 'put',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(para)
    });
    if (response.ok) {
      const responseData = await response.json();
      return responseData;
    } else {
      alert('Failed to get data, please try again later.');
      return null;
    }
  }
  async deleteData(apiUrl = '', para = {}) {
    const response = await fetch(apiUrl, {
      method: 'delete',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(para)
    });
    if (response.ok) {
      const responseData = await response.json();
      return responseData;
    } else {
      alert('Failed to get data, please try again later.');
      return null;
    }
  }
  getFormValue(formTarget = 'form') {
    const form = this.getDom(formTarget, 'form');
    if (!form) {
      return console.warn('getFormValue: parameter(formTarget) is not correct');
    }
    const multipleKey = []; // for multiple select, file
    const formData = new FormData(form);
    const dataObj = {};
    formData.forEach((value, name) => {
      const thisInput = form.querySelector(`[name="${name}"]`);
      // exclude file
      if (thisInput && thisInput.type === 'file') {
        return;
      }
      if (thisInput && thisInput.type === 'number') {
        value = +value;
      }
      if (dataObj[name]) {
        if (!multipleKey.includes(name)) {
          multipleKey.push(name);
          dataObj[name] = [dataObj[name]];
        }
        dataObj[name].push(value);
      } else {
        dataObj[name] = value;
      }
    });
    // array to json
    for (let name in dataObj) {
      if (Array.isArray(dataObj[name]))
        dataObj[name] = JSON.stringify(dataObj[name]);
    }
    return dataObj;
  }
  createColumnByType(type = '', name = '', datasource = ['a', 'b', 'c']) {
    const _this = this;
    const columnSeting = [];
    const columnName = name || `column-${this.getRandom()}`;
    const columnId = `column-id-${this.getRandom()}`;
    switch (type) {
      case 'text':
        columnSeting.push('span', 'New Text', {
          class: 'form-text',
          name: columnName,
          id: columnId
        });
        break;
      case 'input':
      case 'number':
      case 'date':
        columnSeting.push('input', '', {
          type,
          class: 'form-control',
          name: columnName,
          id: columnId
        });
        break;
      case 'textarea':
        columnSeting.push('textarea', '', {
          class: 'form-control',
          name: columnName,
          id: columnId
        });
        break;
      case 'select':
        const options = [];
        datasource.forEach((data) => {
          if (typeof data === 'object') {
            const text = Object.keys(data)[0];
            const value = data[text];
            options.push(new Option(text, value));
          } else {
            options.push(new Option(data));
          }
        });
        columnSeting.push(
          'select',
          '',
          { class: 'form-select', name: columnName, id: columnId },
          options
        );
        break;
      case 'radios':
      case 'checks':
        const radios = [];
        datasource.forEach((data) => {
          const radioId = `${columnName}-${_this.getRandom()}`;
          const radio = _this.createEl(
            'div',
            '',
            {
              class: 'form-check'
            },
            [
              this.createEl('input', '', {
                type: type === 'radios' ? 'radio' : 'checkbox',
                value: data,
                name: columnName,
                id: radioId,
                class: 'form-check-input'
              }),
              this.createEl('label', data, {
                for: radioId,
                class: 'form-check-label'
              })
            ]
          );
          radios.push(radio);
        });
        columnSeting.push(
          'div',
          '',
          {
            class: 'form-check-box'
          },
          radios
        );
        break;
      default:
        this.columnSetting[type] && columnSeting.push(this.columnSetting[type]);
        break;
    }
    return this.createEl(...columnSeting);
  }
  createFloatLabelBox(column, labelName, boxClass = '') {
    const label = this.createEl('label', labelName);
    const floatBox = this.createEl(
      'div',
      '',
      { class: `form-floating ${boxClass}` },
      [column, label]
    );
    return floatBox;
  }
  addColumnBox(type, target = null) {
    if (!target) return;
    const newColumn = this.createColumnByType(type);
    const labelAttrSetting = {
      class: 'form-label',
      for: newColumn.id
    };
    if (type === 'text') delete labelAttrSetting.for;
    const newLabel = this.createEl('label', 'New Column', labelAttrSetting);
    const newColumnBox = this.createEl(
      'div',
      '',
      { class: 'column-box col-12', 'data-type': type },
      [newLabel, newColumn]
    );
    target.append(newColumnBox);
    newColumnBox.click();
  }
  setColValue(target, value = '') {
    const column = this.getDom(target);
    if (!column) {
      console.warn('setColValue: parameter(target) is not correct ');
      return this;
    }
    const form = column.closest('form');
    const name = column.name;
    const type = column.type || column.tagName.toLowerCase();
    switch (type) {
      case 'radio':
      case 'checkbox':
        const oldCheckedInputs = form.querySelectorAll(
          `[name="${name}"]:checked`
        );
        const checkedInput = form.querySelector(
          `[name="${name}"][value="${value}"]`
        );
        for (const oldCheckedInput of oldCheckedInputs) {
          oldCheckedInput.checked = false;
        }
        if (checkedInput) checkedInput.checked = true;
        break;
      case 'span':
        column.textContent = value;
        break;
      default:
        column.value = value;
        break;
    }
    return this;
  }
  toggleSettingByType(columnType = '') {
    const columnSourceSetting = this.panel.querySelector(
      '[name="source-setting"]'
    );
    if (
      columnType === 'radios' ||
      columnType === 'checks' ||
      columnType === 'select'
    ) {
      columnSourceSetting.removeAttribute('disabled');
    } else {
      columnSourceSetting.setAttribute('disabled', 'disabled');
    }
    return this;
  }
  setColumnPanelValue(targetColumnBox = null) {
    this.panel.querySelector('[name="name-setting"]').value = targetColumnBox
      ? targetColumnBox.querySelector('[name]').getAttribute('name')
      : ''; // only getAttribute can get span's name
    this.panel.querySelector('[name="label-setting"]').value = targetColumnBox
      ? targetColumnBox.querySelector('.form-label').innerHTML
      : '';
    this.panel.querySelector('[name="source-setting"]').value = '';
    let classString = '';
    let width = 12;
    if (targetColumnBox) {
      const newType = targetColumnBox.dataset.type;
      this.panel.querySelector('[name="type-setting"]').value = newType;
      this.toggleSettingByType(newType);
      const classList = targetColumnBox.classList;
      classString = classList.value;
      classString = classString
        .replace('column-box', '')
        .replace('editing', '')
        .replace(/col-(1[0-2]|[1-9])/g, '');
      classList.forEach((item) => {
        if (item.indexOf('col-') > -1) width = item.substring(4);
      });
    }
    this.panel.querySelector('[name="width-setting"]').value = targetColumnBox
      ? width
      : '';
    this.panel.querySelector('[name="class-setting"]').value =
      classString.trim();
    this.panel.querySelector('[name="value-setting"]').value = targetColumnBox
      ? this.getColumnOptionValueString(targetColumnBox)
      : '';
    return this;
  }
  getColumnOptionValueString(columnBox) {
    if (!columnBox) return;
    const columnType = columnBox.dataset.type;
    let value = '';
    switch (columnType) {
      case 'input':
      case 'textarea':
      case 'number':
      case 'date':
        value = columnBox.querySelector('.form-control').value;
        break;
      case 'select':
        const options = new Set(columnBox.querySelector('select').options);
        options.forEach((option) => {
          value += `${option.value},`;
        });
        break;
      case 'radios':
      case 'checks':
        const radios = new Set(columnBox.querySelectorAll('input'));
        radios.forEach((radio) => {
          value += `${radio.value},`;
        });
        break;
      case 'text':
        value = columnBox.querySelector('.form-text').innerHTML;
        break;
    }
    return value;
  }
  getRandom() {
    let random = Math.floor(Math.random().toFixed(5) * 1000000);
    if (random < 100000) random = random * 10;
    while (this.randomList.includes(random)) {
      random = Math.floor(Math.random().toFixed(5) * 1000000);
      if (random < 100000) random = random * 10;
    }
    this.randomList.push(random);
    return random;
  }
  changeOptionsBySource(source = '', targetColumnBox = null) {
    if (!targetColumnBox) return;
    if (typeof source === 'string') source = source.split(',');
    const columnType = targetColumnBox.dataset.type;
    if (columnType === 'select') {
      const select = targetColumnBox.querySelector('select');
      select.innerHTML = '';
      source.forEach((option) => {
        select.append(new Option(option));
      });
    } else if (columnType === 'checks' || columnType === 'radios') {
      const name = targetColumnBox.querySelector('.form-check-input').name;
      const oldCheckBox = targetColumnBox.querySelector('.form-check-box');
      oldCheckBox.remove();
      const newColumn = this.createColumnByType(columnType, name, source);
      targetColumnBox.append(newColumn);
    }

    return this;
  }
  setTableValue(table = null, data = []) {
    if (!table || (table && !table.querySelector('tbody'))) return;
    const tbody = table.querySelector('tbody');
    const firstRow = tbody.querySelector('tr');
    tbody.innerHTML = '';
    for (const rowdata of data) {
      const newRow = firstRow.cloneNode(true);
      for (const name in rowdata) {
        const column = newRow.querySelector(`[name="${name}"]`);
        if (!column) continue;
        this.setColValue(column, rowdata[name]);
        tbody.append(newRow);
      }
    }
  }
  setPanelValue(panel = null, data = {}) {
    const allColumn = panel.querySelectorAll('[name]');
    for (const column of allColumn) {
      const name = column.name;
      const colData = data[name];
      this.setColValue(column, colData);
    }
    return this;
  }
  createColumnSettingBox() {
    // first row
    const typeSelect = this.createColumnByType(
      'select',
      'type-setting',
      this.columnType
    );
    const typeSelectBox = this.createFloatLabelBox(
      typeSelect,
      'Column Type',
      'px-1'
    );
    const editLabelInput = this.createColumnByType('input', 'label-setting');
    const editLabelBox = this.createFloatLabelBox(
      editLabelInput,
      'Label Name',
      'px-1'
    );
    const editWidthInput = this.createColumnByType('number', 'width-setting');
    const editWidthBox = this.createFloatLabelBox(
      editWidthInput,
      'Width ( 1 ~ 12 )',
      'px-1'
    );
    const editNameInput = this.createColumnByType('input', 'name-setting');
    const editNameBox = this.createFloatLabelBox(
      editNameInput,
      'Column Name',
      'px-1'
    );
    // second row
    const editValueInput = this.createColumnByType('input', 'value-setting');
    const editValueBox = this.createFloatLabelBox(
      editValueInput,
      'Value',
      'px-1 mt-2'
    );
    const editClassInput = this.createColumnByType('input', 'class-setting');
    const editClassBox = this.createFloatLabelBox(
      editClassInput,
      'Class Name',
      'px-1 mt-2'
    );
    const editSourceInput = this.createColumnByType(
      'select',
      'source-setting',
      ['', { gender: 'male,female' }, { days: 'Mon,Tue,Wed,Thu,Fri,Sat,Sun' }] // test fack data **********
    );
    editSourceInput.setAttribute('disabled', 'disabled');
    const editSourceBox = this.createFloatLabelBox(
      editSourceInput,
      'Source',
      'px-1 mt-2'
    );

    const firstRowBox = this.createEl('div', '', { class: 'd-flex' }, [
      typeSelectBox,
      editNameBox,
      editLabelBox,
      editWidthBox
    ]);
    const secondRowBox = this.createEl('div', '', { class: 'd-flex' }, [
      editValueBox,
      editClassBox,
      editSourceBox
    ]);
    const columnSettingBox = this.createEl('div', '', {}, [
      firstRowBox,
      secondRowBox
    ]);
    return columnSettingBox;
  }
}
const utils = new Core()