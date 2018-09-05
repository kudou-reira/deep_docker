import React, { Component } from 'react';
import OverallStocksComponent from '../components/overallStocksComponent';

class HomeScreen extends Component {
    render() {
        return(
            <div className="container">
                <div className="center">
                    <div>
                        prices
                    </div>
                    <OverallStocksComponent />
                </div>
            </div>
        )
    }
}

export default HomeScreen;