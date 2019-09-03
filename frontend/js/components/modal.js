function modal(){
    var prnt = document.createElement("div");
    prnt.setAttribute('class', "modal fade")
    prnt.setAttribute('id', "updateFieldNamesModal")
    prnt.setAttribute('tabindex', "-1")
    prnt.setAttribute('role', "dialog")
    prnt.setAttribute('aria-labelledby', "exampleModalLabel")
    prnt.setAttribute('aria-hidden', "true");
}

<div class="modal fade" id="updateFieldNamesModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
            <div class="modal-dialog" role="document">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title" id="exampleModalLabel">Edit collection fields</h5>
                  <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                        <div id="selectedCollectionOptions"></div>
                </div>
                <div class="modal-footer">
                  <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                </div>
              </div>
            </div>
          </div>