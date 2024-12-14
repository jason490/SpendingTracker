/* @refresh reload */
import { render } from 'solid-js/web';
import { Route, Router } from '@solidjs/router';

import './index.css';
import App from './App';
import Dashboard from './pages/dashboard/dashboard';

render(() => (
    <Router root={App}>
        <Route path="/" component={Dashboard} />
        <Route path="/add-remove" component={() => <div>hello1</div>} />
        <Route path="/export" component={() => <div>hello3</div>} />
        <Route path="/settings" component={() => <div>hello2</div>} />
    </Router>

), document.getElementById('root'));
