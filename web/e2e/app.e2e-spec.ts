import { AutomatoPage } from './app.po';

describe('ArtooDetoo App', function() {
  let page: AutomatoPage;

  beforeEach(() => {
    page = new AutomatoPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
