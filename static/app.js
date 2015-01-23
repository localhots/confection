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
    var container = document.getElementById('fields');
    for (var i = 0; i < fields.length; i++) {
        var field = fields[i],
            formGroup = makeDivNode('form-group'),
            label = makeLabelNode(field.path, field.title),
            input;

        formGroup.appendChild(label);

        if (field.options !== null) {
            input = makeSelectNode(field.options, !field.is_required)
        } else {
            input = document.createElement('input');
        }

        input.setAttribute('value', field.value);
        input.setAttribute('id', field.path);
        input.setAttribute('class', 'form-control');
        if (field.is_readonly) {
            input.setAttribute('readonly', 'readonly');
        }

        switch (field.type) {
        case "string":
            input.setAttribute('type', 'text');
            break;
        default:
            console.log("Invalid field type: "+ field.type, field)
        }

        formGroup.appendChild(input);
        container.appendChild(formGroup);
    }
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

// TODO: Support for various types
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64
// float32 float64
