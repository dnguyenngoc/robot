import React from 'react';
import './DocumentManagement.scss'


export default class DocumentManagement extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            
        }
        this.handleUploadImage = this.handleUploadImage.bind(this);
    };

    handleUploadImage(ev) {
        ev.preventDefault();
        const data = new FormData();
        data.append('file', this.uploadInput.files[0]);
        data.append('filename', this.fileName.value);
        console.log(data)
        
      }
    
    render() {
        const { rows } = this.state;
        return (
            <form onSubmit={this.handleUploadImage}>
                <div>
                    <input ref={(ref) => { this.uploadInput = ref; }} type="file" />
                </div>
                <br />
                <div>
                    <input ref={(ref) => { this.fileName = ref; }} type="text" placeholder="Enter the desired name of file" />
                </div>
                <br />
                <div>
                    <button>Upload</button>
                </div>
          </form>
        );
    }
        
}   