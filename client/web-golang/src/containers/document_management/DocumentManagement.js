import './DocumentManagement.scss'
import React from 'react';

export default class DocumentManagement extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedFile: new FormData,
            isFileSelect: false,
        }

        this.changeHandler = this.changeHandler.bind(this);
        this.handleSubmission = this.handleSubmission.bind(this);
    };


    changeHandler(event) {
		this.setState({selectedFile: event.target.files[0]});
		this.setState({isFileSelect: true});

	};
    
    handleSubmission(event) {
        event.preventDefault();
        fileService.uploadPnl(this.state.selectedFile);
    }
    

    render() {
        return (
            <div className="main">
                <div className="title">
                     <h3>File Upload</h3>
                </div>
                <input id="file-upload" className="hidden" type="file"></input>
                   
                <label for="file-upload" id="file-drag" class="upload-box">
                    <div id="upload-caption">Drop image here or click to select</div>
                <img id="image-preview" class="hidden" />
                </label>

            </div>
        );
    }
}


 // <form onSubmit={this.handleSubmission}>
            //     <input type="file" name="file" onChange={(event)=> this.changeHandler(event)} />
            //     {this.state.isFileSelect ? <div>
            //         <p>File Name: {this.state.selectedFile.name}</p>
            //         <p>File Type: {this.state.selectedFile.type}</p>
            //         <p>File Size: {this.state.selectedFile.size}</p>
            //     </div> : <p>no File have bean select</p>
            //     }
            //      <input type="submit" value="Submit" />
            // </form>