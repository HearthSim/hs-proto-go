package packages

import (
	"github.com/HearthSim/hs-proto-go/bnet/account_service"
	"github.com/HearthSim/hs-proto-go/bnet/account_types"
	"github.com/HearthSim/hs-proto-go/bnet/attribute"
	"github.com/HearthSim/hs-proto-go/bnet/authentication_service"
	"github.com/HearthSim/hs-proto-go/bnet/challenge_service"
	"github.com/HearthSim/hs-proto-go/bnet/channel_invitation_service"
	"github.com/HearthSim/hs-proto-go/bnet/channel_invitation_types"
	"github.com/HearthSim/hs-proto-go/bnet/channel_owner"
	"github.com/HearthSim/hs-proto-go/bnet/channel_service"
	"github.com/HearthSim/hs-proto-go/bnet/channel_types"
	"github.com/HearthSim/hs-proto-go/bnet/chat_types"
	"github.com/HearthSim/hs-proto-go/bnet/connection_service"
	"github.com/HearthSim/hs-proto-go/bnet/content_handle"
	"github.com/HearthSim/hs-proto-go/bnet/entity"
	"github.com/HearthSim/hs-proto-go/bnet/friends_service"
	"github.com/HearthSim/hs-proto-go/bnet/friends_types"
	"github.com/HearthSim/hs-proto-go/bnet/game_factory"
	"github.com/HearthSim/hs-proto-go/bnet/game_master_service"
	"github.com/HearthSim/hs-proto-go/bnet/game_master_types"
	"github.com/HearthSim/hs-proto-go/bnet/game_utilities_service"
	"github.com/HearthSim/hs-proto-go/bnet/game_utilities_types"
	"github.com/HearthSim/hs-proto-go/bnet/invitation_types"
	"github.com/HearthSim/hs-proto-go/bnet/notification_service"
	"github.com/HearthSim/hs-proto-go/bnet/presence_service"
	"github.com/HearthSim/hs-proto-go/bnet/presence_types"
	"github.com/HearthSim/hs-proto-go/bnet/profanity"
	"github.com/HearthSim/hs-proto-go/bnet/resource_service"
	"github.com/HearthSim/hs-proto-go/bnet/role"
	"github.com/HearthSim/hs-proto-go/bnet/rpc"
	"github.com/HearthSim/hs-proto-go/bnet/rpc_config"
	"github.com/HearthSim/hs-proto-go/bnet/server_pool_types"
	"github.com/HearthSim/hs-proto-go/pegasus/bobnet"
	"github.com/HearthSim/hs-proto-go/pegasus/game"
	"github.com/HearthSim/hs-proto-go/pegasus/shared"
	"github.com/HearthSim/hs-proto-go/pegasus/spectator"
	"github.com/HearthSim/hs-proto-go/pegasus/util"
)

var _ = account_service.GetAccountRequest{}
var _ = account_types.AccountId{}
var _ = attribute.Variant{}
var _ = authentication_service.ModuleLoadRequest{}
var _ = challenge_service.Challenge{}
var _ = channel_invitation_service.AcceptInvitationRequest{}
var _ = channel_invitation_types.ChannelInvitation{}
var _ = channel_owner.GetChannelIdRequest{}
var _ = channel_service.AddMemberRequest{}
var _ = channel_types.Message{}
var _ = chat_types.ChannelState{}
var _ = connection_service.ConnectRequest{}
var _ = content_handle.ContentHandle{}
var _ = entity.EntityId{}
var _ = friends_service.SubscribeToFriendsRequest{}
var _ = friends_types.Friend{}
var _ = game_factory.GameProperties{}
var _ = game_master_service.JoinGameRequest{}
var _ = game_master_types.Player{}
var _ = game_utilities_service.ClientRequest{}
var _ = game_utilities_types.PlayerVariables{}
var _ = invitation_types.Invitation{}
var _ = notification_service.Notification{}
var _ = presence_service.SubscribeRequest{}
var _ = presence_types.RichPresence{}
var _ = profanity.WordFilter{}
var _ = resource_service.ContentHandleRequest{}
var _ = role.Privilege{}
var _ = rpc.NORESPONSE{}
var _ = rpc_config.RPCMethodConfig{}
var _ = server_pool_types.GetLoadRequest{}
var _ = bobnet.PACKETTYPES{}
var _ = game.Tag{}
var _ = shared.CardDef{}
var _ = spectator.JoinInfo{}
var _ = util.ClientOption{}
