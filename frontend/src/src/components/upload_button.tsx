import React from 'react';
import Button from '@material-ui/core/Button';

class UploadButton extends React.Component {
    private onClickHandle = () => {
        alert("onClickHandle");
    };

    public render() {
        return (
            <Button variant="contained" color="primary" onClick={this.onClickHandle}>
                Choose a file.
            </Button>
        );
    }
}

export default UploadButton;
