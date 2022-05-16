# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: isuxportal/resources/team.proto

require 'google/protobuf'

require 'isuxportal/resources/contestant_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("isuxportal/resources/team.proto", :syntax => :proto3) do
    add_message "isuxportal.proto.resources.Team" do
      optional :id, :int64, 1
      optional :name, :string, 2
      optional :leader_id, :int64, 3
      repeated :member_ids, :int64, 4
      optional :final_participation, :bool, 5
      optional :hidden, :bool, 6
      optional :withdrawn, :bool, 7
      optional :disqualified, :bool, 9
      optional :student, :message, 10, "isuxportal.proto.resources.Team.StudentStatus"
      optional :detail, :message, 8, "isuxportal.proto.resources.Team.TeamDetail"
      optional :leader, :message, 16, "isuxportal.proto.resources.Contestant"
      repeated :members, :message, 17, "isuxportal.proto.resources.Contestant"
    end
    add_message "isuxportal.proto.resources.Team.StudentStatus" do
      optional :status, :bool, 1
    end
    add_message "isuxportal.proto.resources.Team.TeamDetail" do
      optional :email_address, :string, 1
      optional :invite_token, :string, 16
    end
  end
end

module Isuxportal
  module Proto
    module Resources
      Team = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.resources.Team").msgclass
      Team::StudentStatus = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.resources.Team.StudentStatus").msgclass
      Team::TeamDetail = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.resources.Team.TeamDetail").msgclass
    end
  end
end
