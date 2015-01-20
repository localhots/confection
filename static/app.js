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
            formGroup = document.createElement('div'),
            label = document.createElement('label'),
            input;

        formGroup.setAttribute('class', 'form-group');
        label.setAttribute('for', field.path);
        label.appendChild(document.createTextNode(field.title));
        formGroup.appendChild(label);

        if (field.options !== null) {
            input = document.createElement('select');

            if (field.is_required === false) {
                var option = document.createElement('option');

                option.setAttribute('value', '');
                option.appendChild(document.createTextNode('Not selected'));
                input.appendChild(option);
            }

            for (var j = 0; j < field.options.length; j++) {
                var value = field.options[j],
                    option = document.createElement('option');

                option.setAttribute('value', value);
                option.appendChild(document.createTextNode(value));
                input.appendChild(option);
            }
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

loadFields(drawForm);

// TODO: Support for various types
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64
// float32 float64
