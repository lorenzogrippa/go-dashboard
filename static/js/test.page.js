

document.addEventListener('DOMContentLoaded', () => {
    initializeFactoryEventListeners();
    initializeAbstractFactoryEventListeners();
    initializeBuilderEventListeners();

});

function initializeBuilderEventListeners() {
    const dogBuilderBtn = document.getElementById("dogBuilderBtn");
    const catBuilderBtn = document.getElementById("catBuilderBtn");
    const dogBuilderOutput = document.getElementById("dog-builder-output");
    const catBuilderOutput = document.getElementById("cat-builder-output");


    dogBuilderOutput.textContent = '';
    catBuilderOutput.textContent = '';

    dogBuilderBtn.addEventListener("click", () => {
        fetch('/api/dog-from-builder', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                dogBuilderOutput.textContent = JSON.stringify(data);
            })
            .catch(error => {
                dogBuilderOutput.textContent = `Error: ${error}`;
            });
    });

    catBuilderBtn.addEventListener("click", () => {
        fetch('/api/cat-from-builder', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                catBuilderOutput.textContent = JSON.stringify(data);
            })
            .catch(error => {
                catBuilderOutput.textContent = `Error: ${error}`;
            });
    });
}

function initializeAbstractFactoryEventListeners() {
    const dogAbstractFactoryBtn = document.getElementById(
      "dogAbstractFactoryBtn"
    );
    const catAbstractFactoryBtn = document.getElementById(
      "catAbstractFactoryBtn"
    );
    const dogAbstractFactoryOutput = document.getElementById(
      "dog-abstract-factory-output"
    );
    const catAbstractFactoryOutput = document.getElementById(
      "cat-abstract-factory-output"
    );


    dogAbstractFactoryOutput.textContent = '';
    catAbstractFactoryOutput.textContent = '';
    dogAbstractFactoryBtn.addEventListener("click", () => { 
        fetch('/api/dog-from-abstract-factory', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                dogAbstractFactoryOutput.textContent = JSON.stringify(data);
            })
            .catch(error => {
                dogAbstractFactoryOutput.textContent = `Error: ${error}`;
            });
    });

    catAbstractFactoryBtn.addEventListener("click", () => {
        fetch('/api/cat-from-abstract-factory', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                catAbstractFactoryOutput.textContent = JSON.stringify(data);
            })
            .catch(error => {
                catAbstractFactoryOutput.textContent = `Error: ${error}`;
            });
    });
}

function initializeFactoryEventListeners() {
    const dogFactoryBtn = document.getElementById("dogFactoryBtn");
    const catFactoryBtn = document.getElementById("catFactoryBtn");
    const dogFactoryOutput = document.getElementById("dog-factory-output");
    const catFactoryOutput = document.getElementById("cat-factory-output");


    dogFactoryOutput.textContent = '';
    catFactoryOutput.textContent = '';

    dogFactoryBtn.addEventListener("click", () => {
        fetch('/api/dog-from-factory', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                dogFactoryOutput.textContent = JSON.stringify(data);
            })
            .catch(error => {
                dogFactoryOutput.textContent = `Error: ${error}`;
            });
    });

    catFactoryBtn.addEventListener("click", () => {
        fetch('/api/cat-from-factory', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                catFactoryOutput.textContent = JSON.stringify(data);
            })
            .catch(error => {
                catFactoryOutput.textContent = `Error: ${error}`;
            });
    });
}

