import React, { Component } from 'react';
import { Form, FormGroup, Label, Input, FormText } from 'reactstrap';

import readXlsxFile from 'read-excel-file'


export default class ExcelUpload extends Component {

    render() {
        return (
            <div>
                <h1>
                    PNL Monthly Management
                </h1>
                <h3>
                    File EXEL Upload
                </h3>
                <div>
                    <input type="file" onChange={this.onFileChange} />
                    <button onClick={this.onFileUpload}>
                        Upload!
                    </button>
                </div>
                {/* {this.fileData()} */}
            </div>
        );
    }
}