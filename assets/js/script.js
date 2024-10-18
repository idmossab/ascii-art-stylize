    // When the button is Clear
    function clearText() {
        document.getElementById('text').value = ""; // Clear the textarea
    }

    // When the form is submitted
    /*document.querySelector('form').onsubmit = function() {
        const textArea = document.getElementById('text'); // Get the textarea element
        const radios = document.querySelectorAll('input[name="banner"]'); // Get all radio buttons
        let isRadioChecked = false; // Initialize a flag to check if any radio is selected

        // Check if the textarea is empty
        if (textArea.value.trim() === "") {
            alert('Please enter text.');
            return false; // Prevent form submission
        }

        // Check if any radio button is checked
        radios.forEach(function(radio) {
            if (radio.checked) {
                isRadioChecked = true; // Set the flag if a radio is checked
            }
        });

        if (!isRadioChecked) {
            alert('Please select a banner.');
            return false; // Prevent form submission
        }

        // Check the length of the text
        if (textArea.value.length > 700) {
            alert('Please enter text less than 700 characters.');
            return false; // Prevent form submission
        }

        // Check for non-ASCII characters
        for (let i = 0; i < textArea.value.length; i++) {
            if (textArea.value.charCodeAt(i) > 127) {
                alert('Please enter ASCII characters only.'); // Alert the user
                return false; // Prevent form submission
            }
        }

        return true; // Allow form submission if all checks pass
    };*/