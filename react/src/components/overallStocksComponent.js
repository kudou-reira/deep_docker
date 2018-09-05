import React, { Component } from 'react';
import *  as actions from '../actions';
import { connect } from 'react-redux';

class OverallStocksComponent extends Component {
    componentDidMount() {
        this.props.fetchDeepAPI();
    }

    render() {
        return(
            <div>
                This is the stocks table
            </div>
        );
    }
}

function mapStateToProps(state) {
    return {
        results: state.deep
    }
}

export default connect(mapStateToProps, actions)(OverallStocksComponent);