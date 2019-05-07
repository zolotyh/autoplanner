const filterNonSerializableInputs = li => li.filter(i => i.name);

document.addEventListener("DOMContentLoaded", function() {
  const valueList = compose(
    getForm,
    getFormElements,
    any2array,
    filterNonSerializableInputs,
    serializeForm
  );

  getForm().addEventListener("submit", e => {
    e.stopPropagation();
    e.preventDefault();
    document.body.append(valueList());
  });
});

function serializeForm(valueList) {
  return (
    "?" +
    valueList.map(i => `${encodeURI(i.name)}=${encodeURI(i.value)}`).join("&")
  );
}

function getForm() {
  return document.querySelector("#mainForm");
}

function getFormElements(form) {
  return form.elements;
}

// concatinate 2 functions
function compose() {
  funcList = any2array(arguments);
  var first = true;
  return function() {
    const args = any2array(arguments);
    return funcList.reduce((memo, value) => {
      var result = first ? value.apply(this, memo) : value(memo);
      first = false;
      return result;
    }, args);
  };
}

function any2array(input) {
  return [].slice.call(input);
}
