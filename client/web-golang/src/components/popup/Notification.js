import React from 'react';
import './Notification.scss'


export default class Notification extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            type: 'info', //custom, 
            name: "Welcome",
            content: "You've got some friends nearby, stop looking at your phone and find them..."
        }
    };
    render(){
        const { type, name, content } = this.state;
        return (
            <section class="main">
                <div class={"alert alert-" + type}>
                <div class="alert-container">
                <div class="alert-icon">
                    <i class="fa fa-info-circle"></i>
                    </div>
                        <button type="button" class="close-icon" data-dismiss="alert" aria-label="Close">
                                            <span>clear</span>
                        </button>
                        <b class="alert-info">{name}:</b> {content}
                    </div>
                </div>


    {/* <div class="alert alert-custom">
        <div class="alert-container">
            <div class="alert-icon">
                <i class="fa fa-warning"></i>
            </div>
            <button type="button" class="close-icon" data-dismiss="alert" aria-label="Close">
                                <span>clear</span>
            </button>
            <b class="alert-info">Custom alert:</b> You've got some error prepare for it...
        </div>
    </div> */}
    {/* <div class="alert alert-success">
        <div class="alert-container">
        <div class="alert-icon">
            <i class="fa fa-check"></i>
        </div>
        <button type="button" class="close-icon" data-dismiss="alert" aria-label="Close">
                            <span>clear</span>
        </button>
        <b class="alert-info">Success alert:</b> Yuhuuu! You've got your $11.99 album from The Weeknd
        </div>
    </div>
    <div class="alert alert-warning">
        <div class="alert-container">
        <div class="alert-icon">
            <i class="fa fa-warning"></i>
        </div>
        <button type="button" class="close-icon" data-dismiss="alert" aria-label="Close">
                            <span>clear</span>
        </button>
        <b class="alert-info">Warning alert:</b> Hey, it looks like you still have the "copyright © 2015" in your footer. Please update it!
        </div>
    </div>
    <div class="alert alert-danger">
        <div class="alert-container">
        <div class="alert-icon">
            <i class="fa fa-info-circle"></i>
        </div>
        <button type="button" class="close-icon" data-dismiss="alert" aria-label="Close">
                            <span>clear</span>
        </button>
        <b class="alert-info">Danger alert:</b> Damn man! You screwed up the server this time. You should find a good excuse for your Boss...
        </div>
    </div> */}
    </section>
        );
    }
}