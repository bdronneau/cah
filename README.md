# CAH - Card Against Humanity

I develop this open source game based on [Cards against Humanity](https://cardsagainsthumanity.com/) during COVID in order to try golang and vuejs together !

## Status
- Game working
- Code does not have unit tests
- I'm using `setInterval` for refresh views, `websocket` will me more efficient
- Tables are in schema `public`
- vuex is not very well implemented

Despite these behaviours, I'm pretty happy because it's working and I can play with my friends.
Moreover, I learn a lot on vuejs development and golang development ...

If I give more love on this project ?
- Use a better structure for Golang API
- Add context logging
- Add prometheus metrics
- Limit the SPA scope to game part and make golang templates for lobby, home, room creation ...
- Include unit tests and cypress

## Usage
- [front](./front/README.md)
- [api](./api/README.md)

## Credits

Sources for cards
- https://github.com/crhallberg/json-against-humanity
- https://github.com/lhoang/Cards-Against-Humanity---Generator


## License

This project is licensed under the MIT license (see LICENSE file).
