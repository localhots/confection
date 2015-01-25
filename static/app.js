/*
 * Confection App
 */

function loadFields(callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/fields.json', true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var fields = JSON.parse(xhr.responseText);
                callback(fields);
            }
        }
    };
    xhr.send(null);
}

function drawForm(fields) {
    var container = document.getElementById('fields'),
        fieldsets = {};

    for (var i = 0; i < fields.length; i++) {
        var field = fields[i];

        if (field.value !== null) {
            var fieldNode = makeFieldNode(field),
                fieldSetName = field.path.match('(.*\/).*')[1];

            if (!fieldsets[fieldSetName]) {
                fieldsets[fieldSetName] = [];
            }
            fieldsets[fieldSetName].push(fieldNode);
        }
    }

    for (var i = 0; i < fields.length; i++) {
        var field = fields[i],
            node;

        if (field.value === null) {
            node = makeFieldSetNode(field.title, fieldsets[field.path + '/']);
        } else if (!field.path.match(/^\/[^\/]+\//)) {
            node = makeFieldNode(field);
        }

        container.appendChild(node);
    }

}

function makeFieldNode(field) {
    var formGroup = makeDivNode('form-group'),
        label = makeLabelNode(field.path, field.title),
        input;

    if (field.options !== null) {
        input = makeSelectNode(field.options, !field.is_required)
    } else {
        input = document.createElement('input');
    }

    if (field.type !== 'bool') {
        input.setAttribute('value', field.value);
        input.setAttribute('class', 'form-control');
    }
    input.setAttribute('id', field.path);
    if (field.is_readonly) {
        input.setAttribute('readonly', 'readonly');
    }

    switch (field.type) {
    case 'string':
        input.setAttribute('type', 'text');
        formGroup.appendChild(label);
        formGroup.appendChild(input);
        break;
    case 'bool':
        input.setAttribute('type', 'checkbox');
        if (field.value) {
            input.setAttribute('checked', 'checked');
        }
        label.innerHTML = '';
        label.appendChild(input);
        label.appendChild(document.createTextNode(field.title));
        formGroup.setAttribute('class', 'checkbox');
        formGroup.appendChild(label);
        break;
    case 'int':
    case 'int8':
    case 'int16':
    case 'int32':
    case 'int64':
    case 'uint':
    case 'uint8':
    case 'uint16':
    case 'uint32':
    case 'uint64':
    case 'float32':
    case 'float64':
        input.setAttribute('type', 'number');
        switch (field.type) {
        case 'int8':
            input.setAttribute('min', '-128');
            input.setAttribute('max', '127');
            break;
        case 'int16':
            input.setAttribute('min', '-32768');
            input.setAttribute('max', '32767');
            break;
        case 'int32':
            input.setAttribute('min', '-2147483648');
            input.setAttribute('max', '2147483647');
            break;
        case 'int': // Assuming x86-64 architecture
        case 'int64':
            input.setAttribute('min', '-9223372036854775808');
            input.setAttribute('max', '9223372036854775807');
            break;
        case 'uint8':
            input.setAttribute('min', '0');
            input.setAttribute('max', '255');
            break;
        case 'uint16':
            input.setAttribute('min', '0');
            input.setAttribute('max', '65535');
            break;
        case 'uint32':
            input.setAttribute('min', '0');
            input.setAttribute('max', '4294967295');
            break;
        case 'uint': // Assuming x86-64 architecture
        case 'uint64':
            input.setAttribute('min', '0');
            input.setAttribute('max', '18446744073709551615');
            break;
        case 'float32':
        case 'float64':
            input.setAttribute('step', 'any');
            break;
        }
        formGroup.appendChild(label);
        formGroup.appendChild(input);
        break;
    default:
        console.log('Invalid field type: '+ field.type, field)
    }

    return formGroup;
}

function makeFieldSetNode(title, fields) {
    var fieldsetNode = document.createElement('fieldset'),
        legendNode = document.createElement('legend');

    legendNode.appendChild(document.createTextNode(title));
    fieldsetNode.appendChild(legendNode);

    for (var i = 0; i < fields.length; i++) {
        fieldsetNode.appendChild(fields[i]);
    }

    return fieldsetNode;
}

function makeDivNode(classes) {
    var div = document.createElement('div');
    div.setAttribute('class', classes);
    return div;
}

function makeLabelNode(forId, text) {
    var label = document.createElement('label'),
        contents = document.createTextNode(text);
    label.setAttribute('for', forId);
    label.appendChild(contents);
    return label;
}

function makeSelectNode(options, hasEmptyOption) {
    var select = document.createElement('select');

    if (hasEmptyOption === true) {
        var option = document.createElement('option');

        option.setAttribute('value', '');
        option.appendChild(document.createTextNode('Not selected'));
        select.appendChild(option);
    }

    for (var i = 0; i < options.length; i++) {
        var value = options[i],
            option = document.createElement('option');

        option.setAttribute('value', value);
        option.appendChild(document.createTextNode(value));
        select.appendChild(option);
    }

    return select;
}

loadFields(drawForm);
